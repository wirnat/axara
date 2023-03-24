package v1

import (
	"bufio"
	"github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
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

	defer fileE.Close()

	//Collect all data from executed model model
	scanner := bufio.NewScanner(fileE)

	for scanner.Scan() {
		line := scanner.Text()
		executorModelFound := strings.Contains(line, "@Register")
		if executorModelFound {
			initiator := strings.Fields(line)
			if len(initiator) != 2 {
				return nil, errors.InvalidModelFlag
			}
			modelName := initiator[1]
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
					modelTrait = NewModelTraitFromFile(file, modelName, c)
					err = file.Close()
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	return
}
