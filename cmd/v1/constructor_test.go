package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parser(t *testing.T) {
	conYaml := &Constructor{
		Key:        "ᬅᬓ᭄ᬱᬭ",
		ModelPath:  "example/model",
		ModuleName: "github.com/wirnat/axara",
		ExecuteModels: []string{
			"Company",
			"Branch",
		},
		ModuleTraits: []ModuleTrait{
			{
				Name:     "~model_snake~ repository interfaces",
				Dir:      "~result_path~/~model_snake~/repository/~model_snake~_repository",
				FileName: "~model_snake~.go",
				Template: "template/clean_architecture/repository_interfaces.text",
				Active:   true,
			},
			{
				Name:     "~model_snake~ store param",
				Dir:      "~result_path~/~model_snake~/request/~model_snake~_request",
				FileName: "~model_snake~_store.go",
				Template: "template/clean_architecture/param_store.text",
				Active:   true,
			},
		},
		Meta: map[string]string{
			"result_path": "module",
		},
		IncludeModuleTraits: []string{
			"testing_env/module_trait.yaml",
		},
		ResultPath: "",
		Models: map[string]map[string]interface{}{
			"User": {
				"module": "user",
			},
			"Company": {
				"module": "company",
			},
		},
	}

	conJSON := &Constructor{
		Key:        "ᬅᬓ᭄ᬱᬭ",
		ModelPath:  "example/model",
		ModuleName: "github.com/wirnat/axara",
		ExecuteModels: []string{
			"Company",
			"Branch",
		},
		ModuleTraits: []ModuleTrait{
			{
				Name:     "~model_snake~ repository interfaces",
				Dir:      "~result_path~/~model_snake~/repository/~model_snake~_repository",
				FileName: "~model_snake~.go",
				Template: "template/clean_architecture/repository_interfaces.text",
				Active:   true,
			},
			{
				Name:     "~model_snake~ store param",
				Dir:      "~result_path~/~model_snake~/request/~model_snake~_request",
				FileName: "~model_snake~_store.go",
				Template: "template/clean_architecture/param_store.text",
				Active:   true,
			},
		},
		Meta: map[string]string{
			"result_path": "module",
		},
		IncludeModuleTraits: []string{
			"testing_env/module_trait.json",
		},
		ResultPath: "",
		Models: map[string]map[string]interface{}{
			"User": {
				"module": "user",
			},
			"Company": {
				"module": "company",
			},
		},
	}

	t.Run("Match to Constructor struct", func(t *testing.T) {
		b, err := NewConstructor("testing_env/test_construct.yaml")
		assert.Nil(t, err)
		assert.Equalf(t, conYaml, b, "fromYAML()")

		c, err := NewConstructor("testing_env/test_construct.json")
		assert.Equalf(t, conJSON, c, "fromJSON()")
		assert.Nil(t, err)
	})
}
