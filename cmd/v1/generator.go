package v1

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/janeczku/go-spinner"
	er "github.com/wirnat/axara/cmd/v1/errors"
	"io/fs"
	"io/ioutil"
	"os"
	"text/template"
)

type generator struct {
	GetModelTrait FileModelTrait
	ReaderMeta    ReaderMeta
	Puller        Puller
}

func NewGenerator(getModelTrait FileModelTrait, readerMeta ReaderMeta, puller Puller) *generator {
	return &generator{GetModelTrait: getModelTrait, ReaderMeta: readerMeta, Puller: puller}
}

/*
	Generate generate file base on declared variable on constructor
*/

var ss = spinner.StartNew("Wait...")

var yesForAll *bool

func (g generator) Generate(c Constructor) error {
	defer ss.Stop()
	ya := false
	yesForAll = &ya

	if c.Key != "ᬅᬓ᭄ᬱᬭ" {
		return er.InvalidKey
	}
	if c.ExecuteModels == nil {
		return er.NothingTodo
	}
	if c.ModuleTraits == nil {
		return er.NothingTodo
	}

	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return er.NoModelFound
	}

	fmt.Println("")

	var mt []*ModelTrait
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

	err = g.generateOnce(c)
	if err != nil {
		return err
	}

	err = g.generatePerModule(mt, mf, c)
	if err != nil {
		return err
	}

	return nil
}

/*
	generateOnce generate file in traits
	this method executed only once when the generator executed
*/
func (g generator) generateOnce(c Constructor) error {
	decoder := NewDecoder(c)
	for _, trait := range c.Traits {
		if !trait.Active {
			continue
		}

		if trait.Remote != "" {
			dirTarget := decoder.Decode(trait.Dir, nil)
			remoteLink := decoder.Decode(trait.Remote, nil)
			f, err := ioutil.ReadDir(dirTarget)
			if len(f) < 1 {
				err = g.Puller.Pull(remoteLink, dirTarget)
				if err != nil {
					return err
				}
			}
		} else {
			//TODO: get builder and set to template
		}
	}

	return nil
}

/*
	generatePerModule generate file per scanned model
	module traits will loop inside scan model process
*/
func (g generator) generatePerModule(mt []*ModelTrait, mf []fs.FileInfo, c Constructor) error {
	totalTask := 0
	successTask := 0
	//loop scanned model trait
	for i, t := range mt {
		//generate file per module model
		for _, trait := range c.ModuleTraits {
			totalTask++
			if !trait.Active {
				continue
			}
			builder := ModuleBuilder{
				Constructor: c,
				ModelTrait:  *t,
			}

			//get meta from model file
			metas, err := g.ReaderMeta.GetMeta(mf[i], c)
			if err != nil {
				return err
			}
			for key, val := range metas {
				builder.Meta[key] = val
			}

			//decode ~code~
			decoderTrait := NewDecoderTrait(builder.Constructor)
			trait = decoderTrait.DecodeTrait(trait, &builder.ModelTrait)

			decoder := NewDecoderBuilder(builder.Constructor)
			builder = decoder.DecodeBuilder(builder)

			err = os.MkdirAll(trait.Dir, os.ModePerm)
			if err != nil {
				fmt.Printf("	❌ create directory failed")
				continue
			}

			tmt, err := template.ParseFiles(trait.Template)
			if err != nil {
				fmt.Printf("	❌ read template failed")
				continue
			}

			generatedFile := fmt.Sprintf("%v/%v", trait.Dir, trait.FileName)

			if !*yesForAll {
				if _, err := os.Stat(generatedFile); !errors.Is(err, os.ErrNotExist) {
					var input string
				Scan:
					{
						ss.Stop()
						fmt.Println(trait.FileName+" is already exist, do you want to override?", "Y=Yes", "N=No", "YA=Yes for all")
						_, err := fmt.Scanln(&input)
						if err != nil {
							fmt.Printf("	something is wrong")
							continue
						}
					}

					input = strcase.ToSnake(input)
					if input == "ya" {
						ya := true
						yesForAll = &ya
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
				fmt.Printf("	❌ create file failed")
				continue
			}

			err = tmt.Execute(fileTrait, builder)
			if err != nil {
				fmt.Printf("	❌ compile template failed")
				continue
			}
			err = fileTrait.Close()
			if err != nil {
				return err
			}
			fmt.Printf("	✅  %v\n", trait.Name)
			successTask++
		}
	}
	if len(mt) < 1 {
		return er.NoModelCanExecute
	}

	fmt.Printf("====== Generate Module Trait Files , %v/%v ======= \n", successTask, totalTask)

	return nil
}
