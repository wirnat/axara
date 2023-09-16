package v1

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/wirnat/axara/infrastructure/ztring"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ModelTrait struct {
	ModelFields []ModelField `json:"model_fields"`
	Model       string       `json:"model"`
	ModelSnake  string       `json:"model_snake"`
	ModelCamel  string       `json:"model_camel"`
	ModelPlural string       `json:"model_plural"`
	ModelHyp    string       `json:"model_hyp"`
	FileInfo    fs.FileInfo
	ModelMeta   map[string]string
}

//getModelField get model fields
func (r *ModelTrait) getModelField(fl io.Reader, config Constructor) error {
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "~ignore") {
			continue
		}

		if strings.Contains(scanner.Text(), fmt.Sprintf("@Register %v", r.Model)) {
			break
		}

		field := strings.Fields(scanner.Text())
		if len(field) == 1 {
			err := r.getEmbeddedModelField(field[0], config)
			if err != nil {
				return err
			}
		}

		if strings.Contains(scanner.Text(), "json") {
			line := strings.Fields(scanner.Text())
			isPtr := strings.Contains(line[1], "*")
			mf := ModelField{
				Json:  strcase.ToSnake(line[0]),
				Name:  line[0],
				Type:  strings.Replace(line[1], "*", "", 1),
				IsPtr: isPtr,
			}
			if strings.Contains(scanner.Text(), "@meta") {
				metas := strings.SplitAfter(scanner.Text(), "@meta")
				metf := strings.Fields(metas[1])
				mf.Meta = map[string]interface{}{}
				for _, m := range metf {
					meta := strings.Split(m, ":")
					mf.Meta[meta[0]] = meta[1]
				}
			}

			r.ModelFields = append(r.ModelFields, mf)
		}

	}
	return nil
}

func (r *ModelTrait) getEmbeddedModelField(modelName string, config Constructor) error {
	modelName = strcase.ToSnake(modelName)
	//TODO: search file name base on model name
	files, err := ioutil.ReadDir(config.ModelPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := strings.Split(file.Name(), ".go")[0]
		if fileName == modelName {
			fileName := filepath.Join(config.ModelPath, file.Name())
			fileOS, err := os.Open(fileName)
			if err != nil {
				return err
			}

			err = r.getModelField(fileOS, config)
			if err != nil {
				return err
			}

			err = fileOS.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewModelTraitFromFile(fl io.Reader, modelName string, config Constructor) *ModelTrait {

	m := &ModelTrait{
		Model:       modelName,
		ModelSnake:  strcase.ToSnake(modelName),
		ModelCamel:  strcase.ToLowerCamel(modelName),
		ModelPlural: ztring.Pluralize(strcase.ToSnake(modelName)),
		ModelHyp:    ztring.ConvertToHyphenated(strcase.ToSnake(modelName)),
	}

	m.getModelField(fl, config)
	return m
}
