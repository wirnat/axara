package decoder

import (
	"fmt"
	"github.com/wirnat/axara/cmd/v1"
	"strings"
)

type decoder struct {
	Construct v1.Constructor
}

func NewDecoder(construct v1.Constructor) *decoder {
	return &decoder{Construct: construct}
}

/*
	Decode parsing the ~~ code
*/
func (d decoder) Decode(code string, mt *v1.ModelTrait) (encoded string) {
	encoded = code
	if mt != nil {
		encoded = strings.ReplaceAll(encoded, "~model_snake~", mt.ModelSnake)
		encoded = strings.ReplaceAll(encoded, "~model~", mt.Model)
		encoded = strings.ReplaceAll(encoded, "~model_camel~", mt.ModelCamel)
	}
	encoded = strings.ReplaceAll(encoded, "~model_path~", d.Construct.ModelPath)
	encoded = strings.ReplaceAll(encoded, "~module_name~", d.Construct.ModuleName)
	encoded = strings.ReplaceAll(encoded, "~model_path~", d.Construct.ModelPath)
	for key, val := range d.Construct.Meta {
		meta := fmt.Sprintf("~%v~", key)
		encoded = strings.ReplaceAll(encoded, meta, val)
	}

	return
}
