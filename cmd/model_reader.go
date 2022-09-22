package cmd

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/wirawirw/aksara-cli/util/stringtor"
	"io"
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

		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			return nil, fmt.Errorf("failed read file " + file.Name())
		}

		//model must contain @model to generate the module
		isExecutedModel := strings.Contains(string(fileContent), "@model")
		if isExecutedModel {
			fileE, err := os.Open(fileName)
			if err != nil {
				return nil, err
			}

			var trait Trait

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
					trait.ModelNameSnake = stringtor.ToSnakeCase(modelName)
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
				bl.Jobs = append(bl.Jobs, trait)
			}

			fileE.Close()
		}

		file.Close()
	}

	return
}

func getCode(file io.Reader, code string) string {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, code) {
			field := strings.Fields(line)
			return field[1]
		}
	}

	return ""
}
