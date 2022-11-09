package v0

import (
	"fmt"
	"strings"
)

type decoder struct {
	*Blueprint
}

func NewDecoder(blueprint *Blueprint) *decoder {
	return &decoder{Blueprint: blueprint}
}

func (d *decoder) Decode() (*Blueprint, error) {
	for i, _ := range d.Blueprint.Jobs {
		j := &d.Blueprint.Jobs[i]

		d.ResultPath = d.convert(d.ResultPath, j)
		d.ModuleName = d.convert(d.ModuleName, j)
		j.Name = d.convert(j.Name, j)
		j.Dir = d.convert(j.Dir, j)
		j.Template = d.convert(j.Template, j)
		j.FileName = d.convert(j.FileName, j)
		j.ModuleTrait.Name = d.convert(j.ModuleTrait.Name, j)
		j.ModuleTrait.Dir = d.convert(j.ModuleTrait.Dir, j)
		j.ModuleTrait.FileName = d.convert(j.ModuleTrait.FileName, j)
		j.ModuleTrait.Template = d.convert(j.ModuleTrait.Template, j)
		for k, val := range j.ModuleTrait.Meta {
			j.Meta[k] = d.convert(val.(string), j)
		}
	}
	return d.Blueprint, nil
}

func (d *decoder) convert(code string, trait *Trait) (r string) {
	code = strings.ReplaceAll(code, "~model_path~", d.ModelPath)

	code = strings.ReplaceAll(code, "~result_path~", d.ResultPath)
	code = strings.ReplaceAll(code, "~model_snake~", trait.ModelSnake)
	code = strings.ReplaceAll(code, "~model~", trait.Model)
	code = strings.ReplaceAll(code, "~model_snake~", trait.ModelCamel)
	code = strings.ReplaceAll(code, "~model_camel~", trait.ModelCamel)
	code = strings.ReplaceAll(code, "~module_name~", d.ModuleName)
	code = strings.ReplaceAll(code, "~model_path~", d.ModelPath)

	for key, val := range trait.Meta {
		newMetaWithoutAd := strings.ReplaceAll(key, "@", "")
		trait.Meta[newMetaWithoutAd] = val
		waterKey := fmt.Sprintf("~%v~", newMetaWithoutAd)
		code = strings.ReplaceAll(code, waterKey, val.(string))
	}

	code = strings.ReplaceAll(code, "~", "")
	return code
}
