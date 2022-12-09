package v1

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	er "github.com/wirnat/axara/cmd/v1/errors"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
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

var yesForAll *bool

func (g generator) Generate(c Constructor) error {
	//defer ss.Stop()
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

	ss.Title = "Read model path..."
	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return er.NoModelFound
	}

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

	//ss.Stop()

	return nil
}

/*
	generateOnce generate file in traits
	this method executed only once when the generator executed
*/
func (g generator) generateOnce(c Constructor) error {
	decoder := NewDecoder(c)
	for _, trait := range c.Traits {
		ss.Title = fmt.Sprintf("Execute %v... ", trait.Name)

		if !trait.Active {
			continue
		}

		if len(trait.CMD) > 1 {
			dirTarget := decoder.Decode(trait.Dir, nil)

			if _, err := os.Stat(dirTarget); errors.Is(err, os.ErrNotExist) {
				cmd := exec.Command(trait.CMD[0], trait.CMD[1:]...)
				err = cmd.Run()
				if err != nil {
					return fmt.Errorf("invalid command")
				}
			}
		} else {
			constructorDecoded := Constructor{
				GitAccessKey:        decoder.Decode(c.GitAccessKey, nil),
				Key:                 decoder.Decode(c.Key, nil),
				ModelPath:           decoder.Decode(c.ModelPath, nil),
				ResultPath:          decoder.Decode(c.ResultPath, nil),
				ModuleName:          decoder.Decode(c.ModuleName, nil),
				ModuleTraits:        nil,
				Meta:                nil,
				IncludeModuleTraits: c.IncludeModuleTraits,
				IncludeTraits:       c.IncludeTraits,
				Traits:              nil,
			}

			for _, moduleTrait := range c.ModuleTraits {
				moduleTrait = ModuleTrait{
					Name:     decoder.Decode(moduleTrait.Name, nil),
					Dir:      decoder.Decode(moduleTrait.Dir, nil),
					FileName: decoder.Decode(moduleTrait.FileName, nil),
					Template: decoder.Decode(moduleTrait.Template, nil),
					Active:   moduleTrait.Active,
					CMD:      moduleTrait.CMD,
				}
				constructorDecoded.ModuleTraits = append(constructorDecoded.ModuleTraits, moduleTrait)
			}

			for key, val := range constructorDecoded.Meta {
				constructorDecoded.Meta[key] = decoder.Decode(val, nil)
			}

			for _, t := range constructorDecoded.Traits {
				t = ModuleTrait{
					Name:     decoder.Decode(t.Name, nil),
					Dir:      decoder.Decode(t.Dir, nil),
					FileName: decoder.Decode(t.FileName, nil),
					Template: decoder.Decode(t.Template, nil),
					Active:   t.Active,
					CMD:      t.CMD,
				}
				constructorDecoded.Traits = append(constructorDecoded.Traits, t)
			}

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
			ss.Title = fmt.Sprintf("Build %v... ", trait.Name)

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
				fmt.Println("	❌ create directory failed")
				continue
			}

			templateDir := decoder.Decode(trait.Template, nil)

			tmt, err := template.ParseFiles(templateDir)
			if err != nil {
				fmt.Println("	❌ read template failed, " + templateDir)
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
							fmt.Println("	something is wrong")
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
				fmt.Println("	❌ create file failed")
				continue
			}

			err = tmt.Execute(fileTrait, builder)
			if err != nil {
				fmt.Println("	❌ compile template failed")
				continue
			}
			err = fileTrait.Close()
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf("	✅  %v \n", trait.Name))
			successTask++
		}
	}
	if len(mt) < 1 {
		return er.NoModelCanExecute
	}

	fmt.Println(fmt.Sprintf("====== Generate Module Trait Files , %v/%v ======= ", successTask, totalTask))

	return nil
}
