package generator

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/wirnat/axara/cmd/v1"
	er "github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type generator struct {
	GetModelTrait v1.ReaderModel
	ReaderMeta    v1.ReaderMeta
	Decoder       v1.Decoder
}

func NewGenerator(getModelTrait v1.ReaderModel, readerMeta v1.ReaderMeta, decoder v1.Decoder) *generator {
	return &generator{GetModelTrait: getModelTrait, ReaderMeta: readerMeta, Decoder: decoder}
}

/*
	Generate generate file base on declared variable on constructor
*/

var yesForAll bool

func (g generator) Generate(c v1.Constructor) error {
	if c.Key != "ᬅᬓ᭄ᬱᬭ" {
		return er.InvalidKey
	}

	if c.Jobs == nil {
		return er.NothingTodo
	}

	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return er.NoModelFound
	}

	var mt []*v1.ModelTrait
	var mf []fs.FileInfo
	//get meta from model and get scanned model trait
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		modelTrait, err := g.GetModelTrait.GetModelTrait(file, c)
		if err != nil {
			return err
		}

		if modelTrait != nil && modelTrait.Model != "" {
			mf = append(mf, file)
			mt = append(mt, modelTrait)
		}
	}

	err = g.executePerModel(mt, mf, c)
	if err != nil {
		return err
	}

	err = g.executeSingle(c)
	if err != nil {
		return err
	}

	return nil
}

func (g generator) getData(c v1.Constructor, t *v1.ModelTrait, mf fs.FileInfo) (builder v1.ModuleBuilder, err error) {
	builder = v1.ModuleBuilder{
		Constructor: c,
		ModelTrait:  t,
	}

	//get meta from model file
	if builder.ModelTrait != nil {
		metas, err := g.ReaderMeta.GetMeta(mf, c, builder.Model)
		if err != nil {
			return v1.ModuleBuilder{}, err
		}

		for key, val := range metas {
			builder.Meta[key] = val
		}

		builder = g.Decoder.DecodeBuilder(builder)

		//decode ~code~
		for i, job := range c.Jobs {
			builder.Constructor.Jobs[i] = g.Decoder.DecodeTrait(job, builder.ModelTrait)
		}
	}
	return
}

func (g generator) executeSingle(constructor v1.Constructor) error {
loop:
	for _, job := range constructor.Jobs {
		if !job.SingleExecute {
			continue loop
		}
		generatedFile := fmt.Sprintf("%v/%v", job.Dir, job.FileName)
		generatedFile = g.Decoder.Decode(generatedFile, nil)

		if !yesForAll {
			if _, err := os.Stat(generatedFile); !errors.Is(err, os.ErrNotExist) {
				var input string
			Scan:
				{
					fmt.Println(job.FileName+" is already exist, do you want to override?", "Y=Yes", "N=No", "YA=Yes for all")
					_, err := fmt.Scanln(&input)
					if err != nil {
						fmt.Println("	something is wrong")
						continue
					}
				}

				input = strcase.ToSnake(input)
				if input == "ya" {
					ya := true
					yesForAll = ya
				}
				if input == "no" || input == "n" {
					continue
				}
				if input != "yes" && input != "y" && input != "ya" {
					goto Scan
				}
			}
		}
		err := os.MkdirAll(g.Decoder.Decode(job.Dir, nil), os.ModePerm)
		if err != nil {
			fmt.Println("	❌" + err.Error())
			continue
		}

		fileTrait, err := os.Create(generatedFile)
		if err != nil {
			return err
		}

		tmt, err := template.ParseFiles(job.Template)
		if err != nil {
			fmt.Println("❌ " + err.Error())
			return err
		}

		if !job.Active {
			continue
		}

		builder, err := g.getData(constructor, nil, nil)
		if err != nil {
			return err
		}

		err = tmt.Execute(fileTrait, builder)
		if err != nil {
			fmt.Println("❌ " + err.Error())
			return err
		}
		err = fileTrait.Close()
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("	✅  %v \n", job.Name))
	}
	return nil
}

/*
	generatePerModule generate file per scanned model
	module traits will loop inside scan model process
*/
func (g generator) executePerModel(mt []*v1.ModelTrait, mf []fs.FileInfo, c v1.Constructor) error {
	totalTask := 0
	successTask := 0
	//loop scanned model trait
	for i, t := range mt {
		//generate file per module model
		builder, err := g.getData(c, t, mf[i])
		if err != nil {
			return err
		}
	loop:
		for _, job := range c.Jobs {
			if job.SingleExecute {
				continue loop
			}
			if global.Tags != nil {
				for _, moduleTrait := range global.Tags {
					if job.Tags == nil {
						continue loop
					}
					for _, tag := range job.Tags {
						if tag != moduleTrait {
							continue loop
						}
					}
				}
			}

			totalTask++
			if !job.Active {
				continue
			}

			err = os.MkdirAll(g.Decoder.Decode(job.Dir, builder.ModelTrait), os.ModePerm)
			if err != nil {
				fmt.Println("	❌" + err.Error())
				continue
			}

			templateDir := g.Decoder.Decode(job.Template, builder.ModelTrait)

			tmt, err := template.ParseFiles(templateDir)
			if err != nil {
				fmt.Println("❌ " + err.Error())
				continue
			}

			generatedFile := fmt.Sprintf("%v/%v", job.Dir, job.FileName)
			generatedFile = g.Decoder.Decode(generatedFile, builder.ModelTrait)
			if job.GenerateIn == "" {
				if !yesForAll {
					if _, err := os.Stat(generatedFile); !errors.Is(err, os.ErrNotExist) {
						var input string
					Scan:
						{
							fmt.Println(job.FileName+" is already exist, do you want to override?", "Y=Yes", "N=No", "YA=Yes for all")
							_, err := fmt.Scanln(&input)
							if err != nil {
								fmt.Println("	something is wrong")
								continue
							}
						}

						input = strcase.ToSnake(input)
						if input == "ya" {
							ya := true
							yesForAll = ya
						}
						if input == "no" || input == "n" {
							continue
						}
						if input != "yes" && input != "y" && input != "ya" {
							goto Scan
						}
					}
				}

				fileTrait, err := os.Create(generatedFile)
				if err != nil {
					fmt.Println("	❌ " + err.Error())
					continue
				}

				err = tmt.Execute(fileTrait, builder)
				if err != nil {
					fmt.Println("	❌ " + err.Error())
					continue
				}
				err = fileTrait.Close()
				if err != nil {
					return err
				}
				fmt.Println(fmt.Sprintf("	✅  %v \n", job.Name))
				successTask++
			} else {
				err := func() error {
					f, err := os.Open(generatedFile)
					if err != nil {
						return err
					}

					defer f.Close()

					// Membaca isi file baris per baris dan menulis kembali ke file yang lama
					tmpFile, err := os.CreateTemp(job.Dir, "test_*.txt")
					if err != nil {
						return err
					}
					defer os.Remove(tmpFile.Name())

					//Collect all data from executed model model
					scanner := bufio.NewScanner(f)
					writer := bufio.NewWriter(tmpFile)

					for scanner.Scan() {
						line := scanner.Text()
						generateCommentFound := strings.Contains(line, "@Generate")
						if generateCommentFound {
							initiator := strings.Fields(line)
							if len(initiator) != 2 {
								return err
							}
							generateTag := initiator[1]
							re, err := regexp.Compile(`[^\w]`)
							if err != nil {
								return err
							}
							generateTag = string(re.ReplaceAll([]byte(generateTag), []byte("")))
							if generateTag == job.GenerateIn {
								var buf strings.Builder
								if err := tmt.Execute(&buf, builder); err != nil {
									return err
								}
								line = buf.String()
							}
						}
						fmt.Fprintln(writer, line)

						err = writer.Flush()
						if err != nil {
							return err
						}
					}

					// Menimpa file lama dengan file yang telah diperbarui
					if err = os.Rename(tmpFile.Name(), generatedFile); err != nil {
						return err
					}
					return nil
				}()
				if err != nil {
					return err
				}
			}

		}
	}
	if len(mt) < 1 {
		return er.NoModelCanExecute
	}

	fmt.Println(fmt.Sprintf("====== Generate Module Trait Files , %v/%v ======= ", successTask, totalTask))

	return nil
}
