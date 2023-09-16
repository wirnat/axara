package runner

import (
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/service/decoder"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2"
	"github.com/wirnat/axara/cmd/v1/service/reader"
)

type app struct {
	path string
}

func NewApp(path string) *app {
	return &app{path: path}
}

func (a app) Generate() error {
	constructorReader := reader.NewReaderConstruct()
	constructor, err := constructorReader.Read(a.path)
	if err != nil {
		return err
	}
	var modelReader v1.ReaderModel
	switch constructor.Lang {
	case v1.Typescript:
		modelReader = reader.NewReaderFileTs()
		break
	case v1.Dart:
		modelReader = reader.NewReaderFileTs()
		break
	default:
		modelReader = reader.NewModelFileReader()
	}

	_decoder := decoder.NewDecoder(constructor)
	readerMeta := reader.NewReaderMeta()
	_generator := generator_v2.NewGenerator(modelReader, _decoder, readerMeta)
	err = _generator.Generate(*constructor)
	if err != nil {
		return err
	}
	return nil
}
