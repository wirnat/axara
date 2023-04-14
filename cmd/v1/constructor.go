package v1

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Constructor struct {
	GitAccessKey  string                       `json:"git_access_key" yaml:"git_access_key"`
	Key           string                       `json:"key" yaml:"key"`
	ModelPath     string                       `json:"model_path" yaml:"model_path"`
	ModuleName    string                       `json:"module_name" yaml:"module_name"`
	Jobs          []Job                        `json:"jobs" yaml:"jobs"`
	Meta          map[string]string            `json:"meta" yaml:"meta"`
	IncludeJobs   []string                     `json:"include_jobs"  yaml:"include_jobs"`
	IncludeTraits []string                     `json:"include_traits" yaml:"include_traits"`
	Models        map[string]map[string]string `json:"models" yaml:"models"`
}

func NewConstructor(p string) (b *Constructor, err error) {
	b, err = marshal(p)
	if err != nil {
		return nil, err
	}

	for _, imt := range b.IncludeJobs {
		c := new(Constructor)
		c, err = marshal(imt)
		if err != nil {
			return nil, err
		}

		b.Jobs = append(b.Jobs, c.Jobs...)
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
