package cmd

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"github.com/wirnat/aksara-cli/util/stringtor"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type modelScanner struct {
	Config
}

func NewModelScanner(c Config) *modelScanner {
	return &modelScanner{Config: c}
}

//Scan  model and get list of trait
func (r modelScanner) Scan() (bl *Blueprint, err error) {
	bl = &Blueprint{}
	bl.ResultPath = r.ResultPath
	bl.ModuleName = r.ModuleName
	bl.ModelPath = r.ModelPath

	files, err := ioutil.ReadDir(r.ModelPath)
	if err != nil {
		return nil, err
	}

	//loop inside model
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		logrus.Info(fmt.Sprintf("Reading %v...", file.Name()))
		fileName := filepath.Join(r.ModelPath, file.Name())
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}

		fileE, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}

		var trait Trait //in the template, trait will be used as a variable

		//Found Custom meta in model
		scanner := bufio.NewScanner(fileE)
		for scanner.Scan() {
			line := scanner.Text()
			//find model name meta
			if strings.Contains(line, "@model") {
				modelName := ""
				field := strings.Fields(line)
				modelName = field[1]

				trait.ModelTrait = *(NewModelTraitFromFile(file, modelName, r.Config))
				trait.ModelSnake = stringtor.ToSnakeCase(modelName)
				trait.ModelCamel = strcase.ToLowerCamel(modelName)
			}
			//find other model meta in model file
			if strings.Contains(line, "@") {
				field := strings.Fields(line)
				if trait.Meta == nil {
					trait.Meta = map[string]interface{}{}
				}
				trait.Meta[field[0]] = field[1]
			}
		}

		for _, bt := range r.ModuleTraits {
			trait.ModuleTrait = bt
			addToJobs := false
			//model must in executed_model in json
			for _, xc := range r.Config.ExecuteModels {
				if xc == trait.Model {
					addToJobs = true
				}
			}
			if addToJobs {
				bl.Jobs = append(bl.Jobs, trait)
			}
		}

		fileE.Close()

	}

	return
}
