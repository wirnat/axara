package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func StructToMap(obj interface{}, tagString string) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		if objType.Field(i).Tag.Get(tagString) != "" {
			data[objType.Field(i).Tag.Get(tagString)] = objVal.Field(i).Interface()
		} else {
			data[objType.Field(i).Name] = objVal.Field(i).Interface()
		}
	}
	return data
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// ParseStructTag get struct tag and extract params
func ParseStructTag(str string, sep string) map[string]string {
	settings := map[string]string{}
	names := strings.Split(str, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}

		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(strings.ToUpper(values[0]))

		if len(values) >= 2 {
			settings[k] = strings.Join(values[1:], ":")
		} else if k != "" {
			settings[k] = k
		}
	}

	return settings
}

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
