package v1

import (
	"fmt"
	"strings"
)

type Decoder interface {
	GetBuilder() ModuleBuilder
	Decode(code string) (encoded string)
}

type decoder struct {
	Builder ModuleBuilder
}

func (d decoder) GetBuilder() ModuleBuilder {
	return d.Builder
}

func NewDecoder(builder ModuleBuilder) *decoder {
	return &decoder{Builder: builder}
}

func (d decoder) Decode(code string) (encoded string) {
	encoded = strings.ReplaceAll(code, "~model_path~", d.Builder.ModelPath)
	encoded = strings.ReplaceAll(encoded, "~model_snake~", d.Builder.ModelSnake)
	encoded = strings.ReplaceAll(encoded, "~model~", d.Builder.Model)
	encoded = strings.ReplaceAll(encoded, "~model_camel~", d.Builder.ModelCamel)
	encoded = strings.ReplaceAll(encoded, "~module_name~", d.Builder.ModuleName)
	encoded = strings.ReplaceAll(encoded, "~model_path~", d.Builder.ModelPath)
	for key, val := range d.Builder.Meta {
		meta := fmt.Sprintf("~%v~", key)
		encoded = strings.ReplaceAll(encoded, meta, val)
	}

	return
}
