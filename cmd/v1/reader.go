package v1

import "io/fs"

type ReaderModel interface {
	GetModelTrait(file fs.FileInfo, c Constructor) (modelTrait *ModelTrait, err error)
}

type ReaderMeta interface {
	GetMeta(file fs.FileInfo, c Constructor, modelName string) (meta map[string]string, err error)
}

type ReaderConstructor interface {
	Read(p string) (b *Constructor, err error)
}
