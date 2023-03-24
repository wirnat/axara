package v1

import (
	"github.com/janeczku/go-spinner"
)

type app struct {
	path string
}

func NewApp(path string) *app {
	return &app{path: path}
}

var ss = spinner.StartNew("Wait...")

func (a app) Generate() error {
	ss.Title = "Parse orchestrator... "
	constructor, err := NewConstructor(a.path)
	if err != nil {
		return err
	}

	modelReader := NewModelFileReader()
	metaReader := NewReaderMeta()
	generator := NewGenerator(modelReader, metaReader)
	err = generator.Generate(*constructor)
	if err != nil {
		return err
	}
	return nil
}
