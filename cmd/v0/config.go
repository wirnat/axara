package v0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Key           string        `json:"key"`
	ModelPath     string        `json:"model_path"`
	ResultPath    string        `json:"result_path"`
	ModuleName    string        `json:"module_name"`
	ExecuteModels []string      `json:"execute_models"`
	ModuleTraits  []ModuleTrait `json:"module_traits"`
}

func OpenConfig(p string) (b Config, err error) {
	targetConfig := fmt.Sprintf("%v", p)
	jsonData, errFile := os.Open(targetConfig)
	if errFile != nil {
		return Config{}, errFile
	}

	defer jsonData.Close()

	byteValue, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(byteValue, &b)
	if err != nil {
		return Config{}, err
	}

	return
}
