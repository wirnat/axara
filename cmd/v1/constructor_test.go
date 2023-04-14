package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parser(t *testing.T) {
	expectedYaml := &Constructor{GitAccessKey: "", Key: "ᬅᬓ᭄ᬱᬭ", ModelPath: "example/model", ModuleName: "github.com/wirnat/axara", Jobs: []Job{Job{Name: "~model_snake~ repository interfaces", Dir: "~result_path~/~model_snake~/repository/~model_snake~_repository", FileName: "~model_snake~.go", Template: "template/clean_architecture/repository_interfaces.text", Active: true, CMD: []string(nil), Tags: []string{"repo"}, GenerateIn: "route", SingleExecute: false}}, Meta: map[string]string{"result_path": "module"}, IncludeJobs: []string{"spam/testing_env/module_trait.yaml"}, IncludeTraits: []string(nil), Models: map[string]map[string]string{"Company": map[string]string{"module": "company"}, "User": map[string]string{"module": "user"}}}
	expectedJSON := &Constructor{GitAccessKey: "", Key: "ᬅᬓ᭄ᬱᬭ", ModelPath: "example/model", ModuleName: "github.com/wirnat/axara", Jobs: []Job{Job{Name: "~model_snake~ repository interfaces", Dir: "~result_path~/~model_snake~/repository/~model_snake~_repository", FileName: "~model_snake~.go", Template: "template/clean_architecture/repository_interfaces.text", Active: true, CMD: []string(nil), Tags: []string(nil)}}, Meta: map[string]string{"result_path": "module"}, IncludeJobs: []string{"spam/testing_env/module_trait.json"}, IncludeTraits: []string(nil), Models: map[string]map[string]string{"Company": map[string]string{"module": "company"}, "User": map[string]string{"module": "user"}}}

	t.Run("Match to Constructor struct", func(t *testing.T) {
		b, err := NewConstructor("spam/testing_env/test_construct.yaml")
		assert.Nil(t, err)
		assert.Equalf(t, expectedYaml, b, "fromYAML()")

		c, err := NewConstructor("spam/testing_env/test_construct.json")
		assert.Equalf(t, expectedJSON, c, "fromJSON()")
		assert.Nil(t, err)
	})
}
