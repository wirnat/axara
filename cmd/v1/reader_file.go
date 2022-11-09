package v1

import (
	"bufio"
	"fmt"
	"github.com/wirnat/axara/cmd/v1/errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileModelTrait interface {
	GetModelTrait(file fs.FileInfo, c Constructor) (modelTrait *ModelTrait, err error)
}

type file struct{}

func NewModelFileReader() *file {
	return &file{}
}

func (g file) GetModelTrait(file fs.FileInfo, c Constructor) (modelTrait *ModelTrait, err error) {
	fileName := filepath.Join(c.ModelPath, file.Name())
	fileE, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	//Collect all data from executed model model
	scanner := bufio.NewScanner(fileE)
	modelOpen := false
	modelClose := false

	for scanner.Scan() {
		line := scanner.Text()
		executorModelFound := strings.Contains(line, "~model~")
		endModelFound := strings.Contains(line, "~end_model~")
		if endModelFound {
			modelClose = true
			continue
		}
		//set ModelTrait when find ~model~ in line of the file
		if executorModelFound {
			modelOpen = true
			modelFields := strings.Fields(line)
			if len(modelFields) != 2 {
				fmt.Println("invalid tag")
				return nil, errors.InvalidModelFlag
			}
			modelName := modelFields[1]

			for _, modelExecuteJson := range c.ExecuteModels {
				re, err := regexp.Compile(`[^\w]`)
				if err != nil {
					return nil, err
				}
				modelName = string(re.ReplaceAll([]byte(modelName), []byte("")))
				if modelExecuteJson == modelName {
					file, err := os.Open(fileName)
					if err != nil {
						return nil, err
					}
					modelTrait = NewModelTraitFromFile(file, modelName, c)
					err = file.Close()
					if err != nil {
						return nil, err
					}
				}
			}
		}

	}

	if modelOpen && !modelClose {
		return nil, errors.NoEndModelFound
	}

	err = fileE.Close()
	if err != nil {
		return nil, err
	}
	return
}
