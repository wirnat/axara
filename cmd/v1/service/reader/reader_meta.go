package reader

import (
	"bufio"
	"fmt"
	"github.com/wirnat/axara/cmd/v1"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type readerMeta struct{}

func NewReaderMeta() *readerMeta {
	return &readerMeta{}
}

func (m readerMeta) GetMeta(file fs.FileInfo, c v1.Constructor, modelName string) (meta map[string]string, err error) {
	if meta == nil {
		meta = make(map[string]string)
	}
	if file == nil {
		return nil, fmt.Errorf("file is invalid")
	}
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
		if strings.Contains(line, "~") {
			field := strings.Fields(line)
			if len(field) == 2 {
				re, err := regexp.Compile(`[^\w]`)
				if err != nil {
					return nil, err
				}
				field[0] = string(re.ReplaceAll([]byte(field[0]), []byte("")))
				field[1] = string(re.ReplaceAll([]byte(field[1]), []byte("")))
				meta[field[0]] = field[1]
			}
		}
	}

	////Collect meta from config
	for modelConf, modelConfMeta := range c.Models {
		if modelName == modelConf {
			for key, insideModelMeta := range modelConfMeta {
				meta[key] = insideModelMeta
			}
		}
	}
	return
}
