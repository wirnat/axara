package decoder

import (
	"fmt"
	"github.com/wirnat/axara/cmd/v1"
	"strings"
)

type decoder struct {
	Construct *v1.Constructor
}

func NewDecoder(construct *v1.Constructor) *decoder {
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
	d.Construct.Meta = decodeToMap(d.Construct.Meta)
	for key, val := range d.Construct.Meta {
		meta := fmt.Sprintf("~%v~", key)
		encoded = strings.ReplaceAll(encoded, meta, val)
	}

	if mt != nil {
		for key, value := range mt.ModelMeta {
			encoded = strings.ReplaceAll(encoded, fmt.Sprintf("~%v~", key), value)
		}
	}

	return
}
func decodeToMap(dataMap map[string]string) map[string]string {
	decodedMap := make(map[string]string)
	for key, value := range dataMap {
		decodedValue := value
		for subKey, subValue := range dataMap {
			decodedValue = strings.Replace(decodedValue, "~"+subKey+"~", subValue, -1)

		}
		decodedMap[key] = decodedValue
	}

	return decodedMap
}

func checkIfMetaFromModels(mapKey string, data map[string]map[string]string) (hasModule bool) {
	// Melakukan iterasi terhadap map pertama
	for _, v := range data {
		// Melakukan iterasi terhadap map kedua
		for key, _ := range v {
			if key == mapKey { // Jika ditemukan key "module"
				return true // Keluar dari loop
			}
		}
	}
	return false
}
