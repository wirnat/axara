package v1

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Constructor struct {
	GitAccessKey        string            `json:"git_access_key" yaml:"git_access_key"`
	Key                 string            `json:"key" yaml:"key"`
	ModelPath           string            `json:"model_path" yaml:"model_path"`
	ResultPath          string            `json:"result_path" yaml:"result_path"`
	ModuleName          string            `json:"module_name" yaml:"module_name"`
	ExecuteModels       []string          `json:"execute_models" yaml:"execute_models"`
	ModuleTraits        []ModuleTrait     `json:"module_traits" yaml:"module_traits"`
	Meta                map[string]string `json:"meta" yaml:"meta"`
	IncludeModuleTraits []string          `json:"include_module_traits"  yaml:"include_module_traits"`
	IncludeTraits       []string          `json:"include_traits" yaml:"include_traits"`
	Traits              []ModuleTrait     `json:"traits" yaml:"traits"`
}

func NewConstructor(p string) (b *Constructor, err error) {
	b, err = marshal(p)
	if err != nil {
		return nil, err
	}

	for _, imt := range b.IncludeModuleTraits {
		c := new(Constructor)
		c, err = marshal(imt)
		if err != nil {
			return nil, err
		}

		b.ModuleTraits = append(b.ModuleTraits, c.ModuleTraits...)
	}
	return
}

//marshal get data from include yaml file
func marshal(p string) (c *Constructor, err error) {
	targetConfig := fmt.Sprintf("%v", p)
	data, errFile := os.Open(targetConfig)
	if errFile != nil {
		return &Constructor{}, errFile
	}
	defer data.Close()

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		return &Constructor{}, errFile
	}
	isJson := strings.Contains(p, ".json")
	if isJson {
		err = yaml.Unmarshal(byteValue, &c)
		if err != nil {
			return nil, err
		}
	} else {
		err = yaml.Unmarshal(byteValue, &c)
		if err != nil {
			return nil, err
		}
	}
	return
}
