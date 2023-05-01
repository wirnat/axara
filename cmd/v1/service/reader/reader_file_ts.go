package reader

import (
	"bufio"
	"fmt"
	plural "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type readerFileTs struct{}

func NewReaderFileTs() *readerFileTs {
	return &readerFileTs{}
}

func (r readerFileTs) GetModelTrait(file fs.FileInfo, c v1.Constructor) (modelTrait *v1.ModelTrait, err error) {
	//Dapatkan file
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
					modelTrait, err = r.getMTFromTSFile(file, modelName)
					if err != nil {
						return nil, err
					}

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

func (r readerFileTs) getMTFromTSFile(file *os.File, modelName string) (m *v1.ModelTrait, err error) {
	m = &v1.ModelTrait{
		Model:       modelName,
		ModelSnake:  strcase.ToSnake(modelName),
		ModelCamel:  strcase.ToLowerCamel(modelName),
		ModelPlural: strings.ToLower(plural.NewClient().Plural(modelName)),
	}

	scanner := bufio.NewScanner(file)
	openScanField := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "~ignore") {
			continue
		}

		if strings.Contains(scanner.Text(), fmt.Sprintf("@Register %v", modelName)) {
			break
		}
		line := scanner.Text()
		if strings.Contains(line, "constructor(") {
			openScanField = true
			continue
		}
		if strings.Contains(line, ")") && openScanField {
			openScanField = false
		}

		if openScanField {
			line := strings.Fields(scanner.Text())
			isPtr := strings.Contains(line[1], "*")
			mf := v1.ModelField{
				Json:  removeSpecialChar(strcase.ToSnake(line[0])),
				Name:  removeSpecialChar(line[0]),
				Type:  removeSpecialChar(line[1]),
				IsPtr: isPtr,
				Meta:  map[string]interface{}{},
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
			m.ModelFields = append(m.ModelFields, mf)
		}
	}
	return
}

func removeSpecialChar(inputString string) string {
	// Ekspresi reguler untuk menghilangkan koma, titik dua, dan titik
	regexPattern := "[,:.*]"
	re := regexp.MustCompile(regexPattern)
	return re.ReplaceAllString(inputString, "")
}
