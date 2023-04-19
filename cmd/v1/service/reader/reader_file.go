package reader

import (
	"bufio"
	"github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type file struct{}

func NewModelFileReader() *file {
	return &file{}
}

func (g file) GetModelTrait(file fs.FileInfo, c v1.Constructor) (modelTrait *v1.ModelTrait, err error) {
	fileName := filepath.Join(c.ModelPath, file.Name())
	fileE, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer fileE.Close()

	//Collect all data from executed model model
	scanner := bufio.NewScanner(fileE)
	var modelName string

	for scanner.Scan() {
		line := scanner.Text()
		executorModelFound := strings.Contains(line, "@Register")
		if executorModelFound {
			initiator := strings.Fields(line)
			if len(initiator) != 2 {
				return nil, errors.InvalidModelFlag
			}
			modelName = initiator[1]
			re, err := regexp.Compile(`[^\w]`)
			if err != nil {
				return nil, err
			}
			modelName = string(re.ReplaceAll([]byte(modelName), []byte("")))
			for _, executeModel := range global.ExecuteModels {
				if executeModel == modelName {
					file, err := os.Open(fileName)
					if err != nil {
						return nil, err
					}
					modelTrait = v1.NewModelTraitFromFile(file, modelName, c)
					err = file.Close()
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	if modelTrait != nil {
		if modelTrait.ModelMeta == nil {
			modelTrait.ModelMeta = make(map[string]string)
		}
	}

	//Collect meta from config
	for modelConf, modelConfMeta := range c.Models {
		if modelName == modelConf {
			for key, insideModelMeta := range modelConfMeta {
				if modelTrait != nil {
					modelTrait.ModelMeta[key] = insideModelMeta
				}
			}
		}
	}

	return
}
