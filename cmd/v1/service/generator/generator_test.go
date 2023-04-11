package generator

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1"
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
		name    string
		args    args
		wantErr error
		init    func()
		fun     func(t assert.TestingT)
	}{
		{
			name: "total generated directory must equal with execute trait on the model",
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
							Name:     "test",
							Dir:      "../../spam/testing_env/modules/t1",
							FileName: "test.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
						{
							Name:     "test 2",
							Dir:      "../../spam/testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("../../spam/testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 2, len(f)) {
					os.RemoveAll("../../spam/testing_env/modules")
				}
			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
			},
		},
		{
			name: "only generate active trait",
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
							Name:     "test",
							Dir:      "../../spam/testing_env/modules/t1",
							FileName: "test.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
						{
							Name:     "test 2",
							Dir:      "../../spam/testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   false,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("../../spam/testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 1, len(f)) {
					os.RemoveAll("../../spam/testing_env/modules")
				}
			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
			},
		},
		{
			name: "only generate tagged trait",
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
							Name:     "test",
							Dir:      "../../spam/testing_env/modules/t1",
							FileName: "test.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group1",
							},
						},
						{
							Name:     "test 2",
							Dir:      "../../spam/testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group1",
							},
						},
						{
							Name:     "test 3",
							Dir:      "../../spam/testing_env/modules/t3",
							FileName: "test_3.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group2",
							},
						},
						{
							Name:     "test 4",
							Dir:      "../../spam/testing_env/modules/t4",
							FileName: "test_4.go",
							Template: "../../spam/testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group1",
							},
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("../../spam/testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 3, len(f)) {
				}
				os.RemoveAll("../../spam/testing_env/modules")

			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
				global.Tags = []string{"group1"}
			},
		},
		{
			name: "inject code in generate tag",
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
			wantErr: nil,
			fun: func(t assert.TestingT) {
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
			init: func() {
				global.ExecuteModels = []string{"Company", "Branch"}
			},
		},
		{
			name: "Exec single job",
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
							Name:          "main",
							Dir:           "../../spam/testing_env/app",
							FileName:      "main.go",
							Template:      "../../spam/testing_env/templates/main.text",
							Active:        true,
							GenerateIn:    "route",
							SingleExecute: true,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("../../spam/testing_env/app")
				if assert.Nil(t, err) && assert.Equal(t, 1, len(f)) {
					os.RemoveAll("../../spam/testing_env/app")
				}
			},
			init: func() {
				global.ExecuteModels = []string{"Company", "Branch"}
				yesForAll = true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				GetModelTrait: reader.NewModelFileReader(),
				ReaderMeta:    reader.NewReaderMeta(),
				Decoder:       decoder.NewDecoder(tt.args.c),
			}
			tt.init()
			err := g.Generate(tt.args.c)
			assert.Equal(t, tt.wantErr, err)
			if tt.fun != nil {
				tt.fun(t)
			}
		})
	}
}
