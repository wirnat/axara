package generator_v2

import (
	"github.com/stretchr/testify/assert"
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/global"
	"github.com/wirnat/axara/cmd/v1/service/decoder"
	"github.com/wirnat/axara/cmd/v1/service/reader"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_generator_Generate(t *testing.T) {

	type args struct {
		c v1.Constructor
	}
	tests := []struct {
		name  string
		args  args
		init  func(t2 *testing.T)
		check func(t2 *testing.T, err error)
	}{
		//Generate singleExecute job"
		{
			name: "Generate singleExecute job",
			args: args{
				c: v1.Constructor{
					Meta: map[string]string{
						"test":     "this is a meta",
						"template": "../../spam/testing_env/templates",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "../../spam/testing_env/model",
					ModuleName: "test 1",
					Jobs: []v1.Job{
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv1",
							FileName:      "main.go",
							Template:      "~template~/main.text",
							Active:        true,
							SingleExecute: true,
						},
					},
				},
			},
			init: func(t2 *testing.T) {

			},
			check: func(t2 *testing.T, err error) {
				f, err := ioutil.ReadDir("../../spam/testing_env/gv1")
				if assert.Nil(t, err) && assert.Equal(t, 1, len(f)) {
					os.RemoveAll("../../spam/testing_env/gv1")
				}
			},
		},
		//Generate only active job
		{
			name: "Generate only active job",
			args: args{
				c: v1.Constructor{
					Meta: map[string]string{
						"test":     "this is a meta",
						"template": "../../spam/testing_env/templates",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "../../spam/testing_env/model",
					ModuleName: "test 2",
					Jobs: []v1.Job{
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv2",
							FileName:      "main.go",
							Template:      "~template~/main.text",
							Active:        false,
							GenerateIn:    "route",
							SingleExecute: true,
						},
					},
				},
			},
			init: func(t2 *testing.T) {

			},
			check: func(t2 *testing.T, err error) {
				_, err = ioutil.ReadDir("../../spam/testing_env/gv2")
				assert.NotNil(t2, err)
			},
		},
		//Generate base on TAG
		{
			name: "Generate base on TAG",
			args: args{
				c: v1.Constructor{
					Meta: map[string]string{
						"test":     "this is a meta",
						"template": "../../spam/testing_env/templates",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "../../spam/testing_env/model",
					ModuleName: "test 1",
					Jobs: []v1.Job{
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv3",
							FileName:      "main.go",
							Template:      "~template~/main.text",
							Active:        true,
							SingleExecute: true,
							Tags:          []string{"A"},
						},
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv3",
							FileName:      "main2.go",
							Template:      "~template~/main.text",
							Active:        true,
							SingleExecute: true,
							Tags:          []string{"A", "B"},
						},
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv3",
							FileName:      "main3.go",
							Template:      "~template~/main.text",
							Active:        true,
							SingleExecute: true,
							Tags:          []string{"B", "A"},
						},
						{
							Name:          "main",
							Dir:           "../../spam/testing_env/gv3",
							FileName:      "main4.go",
							Template:      "~template~/main.text",
							Active:        true,
							SingleExecute: true,
						},
					},
				},
			},
			init: func(t2 *testing.T) {
				global.Tags = []string{
					"A",
				}
			},
			check: func(t2 *testing.T, err error) {
				f, err := ioutil.ReadDir("../../spam/testing_env/gv3")
				if assert.Nil(t, err) && assert.Equal(t, 3, len(f)) {
					os.RemoveAll("../../spam/testing_env/gv3")
				}
			},
		},
		//Inject code to file
		{
			name: "Inject code to file",
			args: args{
				c: v1.Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "../../spam/testing_env/model",
					ModuleName: "~request~",
					Jobs: []v1.Job{
						{
							Name:       "test",
							Dir:        "../../spam/testing_env",
							FileName:   "route.go",
							Template:   "../../spam/testing_env/templates/route.text",
							Active:     true,
							GenerateIn: "route",
						},
					},
				},
			},
			init: func(t2 *testing.T) {
				global.ExecuteModels = []string{"Company", "Branch"}
			},
			check: func(t2 *testing.T, err error) {
				f, err := os.Open("../../spam/testing_env/route.go")
				if err != nil {
					panic(err)
				}

				// Read the contents of the file into a byte slice
				fileContents, err := ioutil.ReadAll(f)

				// Convert the byte slice to a string and search for the text string
				fileContentsStr := string(fileContents)
				isContain := strings.Contains(fileContentsStr, "//Company Route\ncompany := e.Group(\"company\")\ncompany.GET(\"\", func(c echo.Context) error {\n\treturn nil\n})\ncompany.GET(\"\", func(c echo.Context) error {\n\treturn nil\n})\n//Branch Route\nbranch := e.Group(\"branch\")\nbranch.GET(\"\", func(c echo.Context) error {\n\treturn nil\n})\nbranch.GET(\"\", func(c echo.Context) error {\n\treturn nil\n})")
				assert.Equal(t, true, isContain)
			},
		},
		//Generate perModelExecute job
		{
			name: "Generate perModelExecute job",
			args: args{
				c: v1.Constructor{
					Meta: map[string]string{
						"test":        "this is a ~model_snake~",
						"template":    "../../spam/testing_env/templates",
						"result_path": "../../spam/testing_env/modules",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "../../spam/testing_env/model",
					ModuleName: "test 1",
					Jobs: []v1.Job{
						{
							Name:          "~model_snake~ usecase interface",
							Dir:           "~result_path~/~module~",
							FileName:      "~model_snake~.go",
							Template:      "~template~/usecase_interfaces.text",
							Active:        true,
							SingleExecute: false,
						},
					},
					Models: map[string]map[string]string{
						"Company": {
							"module": "company",
						},
						"Branch": {
							"module": "branch",
						},
					},
				},
			},
			init: func(t2 *testing.T) {
				global.Tags = nil
			},
			check: func(t2 *testing.T, err error) {
				f, err := ioutil.ReadDir("../../spam/testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 2, len(f)) {
					os.RemoveAll("../../spam/testing_env/modules")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				ReaderModel: reader.NewModelFileReader(),
				Decoder:     decoder.NewDecoder(&tt.args.c),
				ReaderMeta:  reader.NewReaderMeta(),
			}
			tt.init(t)

			err := g.Generate(tt.args.c)
			tt.check(t, err)
		})
	}
}

func Test_FullGenerate(t *testing.T) {
	tests := []struct {
		name string
		res  func(t2 *testing.T, err error)
		init func(t *testing.T)
	}{
		{
			name: "Test Generate",
			res: func(t2 *testing.T, err error) {
				global.Tags = []string{"Company", "Branch"}
			},
			init: func(t *testing.T) {
				global.Tags = nil
				global.ExecuteModels = []string{"Branch", "Company"}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := v1.NewConstructor("test/unclebob/uncle_bob.yaml")
			if err != nil {
				panic(err)
			}
			r := reader.NewModelFileReader()
			d := decoder.NewDecoder(c)
			rm := reader.NewReaderMeta()
			g := NewGenerator(r, d, rm)
			tt.init(t)
			err = g.Generate(*c)
			tt.res(t, err)
		})
	}
}
