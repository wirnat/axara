package v1

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ReaderMeta interface {
	GetMeta(file fs.FileInfo, c Constructor) (meta map[string]string, err error)
}

type readerMeta struct{}

func NewReaderMeta() *readerMeta {
	return &readerMeta{}
}

func (m readerMeta) GetMeta(file fs.FileInfo, c Constructor) (meta map[string]string, err error) {
	if file == nil {
		return nil, fmt.Errorf("file is invalid")
	}
	fileName := filepath.Join(c.ModelPath, file.Name())
	fileE, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	//Collect all data from executed model model
	scanner := bufio.NewScanner(fileE)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "~") {
			field := strings.Fields(line)
			if meta == nil {
				meta = map[string]string{}
			}
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

	return
}
