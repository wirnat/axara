package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"text/template"
)

type Base struct {
	ModuleName string `json:"module_name"`
	ResultPath string `json:"result_path"`
	ModelPath  string `json:"model_path"`
}

type BaseTrait struct {
	Base
	Trait
}

type Blueprint struct {
	Base
	Jobs []Trait `json:"jobs"`
}

func (b Blueprint) Generate() error {
	for _, j := range b.Jobs {
		var baseTrait BaseTrait
		baseTrait.Trait = j
		baseTrait.Base = b.Base

		err := os.MkdirAll(j.Dir, os.ModePerm)
		if err != nil {
			return err
		}

		logrus.Info(fmt.Sprintf("Build %v...", j.ModuleTrait.Name))
		tmt, err := template.ParseFiles(j.Template)
		if err != nil {
			return err
		}

		generatedFile := fmt.Sprintf("%v/%v", j.ModuleTrait.Dir, j.FileName)
		fileTrait, err := os.Create(generatedFile)
		if err != nil {
			return err
		}

		err = tmt.Execute(fileTrait, baseTrait)
		if err != nil {
			return err
		}
	}
	return nil
}
