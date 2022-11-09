package v1

import (
	"fmt"
	"github.com/wirnat/axara/cmd/v1/errors"
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
func (g generator) Generate(c Constructor) error {
	if c.Key != "ᬅᬓ᭄ᬱᬭ" {
		return errors.InvalidKey
	}
	if c.ExecuteModels == nil {
		return errors.NothingTodo
	}
	if c.ModuleTraits == nil {
		return errors.NothingTodo
	}

	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return errors.NoModelFound
	}

	var mt []*ModelTrait
	var mf []fs.FileInfo
	//get meta from model and get scanned model trait
	fmt.Println("Collect meta from scanned model...")
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
	for _, trait := range c.Traits {
		if !trait.Active {
			continue
		}
		if trait.Remote != "" {
			f, err := ioutil.ReadDir(trait.Dir)
			if len(f) < 1 {
				err = g.Puller.Pull(trait.Remote, trait.Dir)
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
			decoderTrait := NewDecoderTrait(builder)
			trait = decoderTrait.DecodeTrait(trait)

			decoder := NewDecoderBuilder(builder)
			builder = decoder.DecodeBuilder()

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
		return errors.NoModelCanExecute
	}

	fmt.Printf("====== Generate Module Trait Files , %v/%v ======= \n", successTask, totalTask)

	return nil
}
