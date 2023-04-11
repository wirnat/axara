package v1

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1/global"
	"github.com/wirnat/axara/infrastructure/utils"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_generator_Generate(t *testing.T) {
	type args struct {
		c Constructor
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
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					Jobs: []Job{
						{
							Name:     "test",
							Dir:      "testing_env/modules/t1",
							FileName: "test.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
						{
							Name:     "test 2",
							Dir:      "testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 2, len(f)) {
					os.RemoveAll("testing_env/modules")
				}
			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
			},
		},
		{
			name: "only generate active trait",
			args: args{
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					Jobs: []Job{
						{
							Name:     "test",
							Dir:      "testing_env/modules/t1",
							FileName: "test.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
						},
						{
							Name:     "test 2",
							Dir:      "testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   false,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 1, len(f)) {
					os.RemoveAll("testing_env/modules")
				}
			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
			},
		},
		{
			name: "only generate tagged trait",
			args: args{
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					Jobs: []Job{
						{
							Name:     "test",
							Dir:      "testing_env/modules/t1",
							FileName: "test.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group1",
							},
						},
						{
							Name:     "test 2",
							Dir:      "testing_env/modules/t2",
							FileName: "test_2.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group1",
							},
						},
						{
							Name:     "test 3",
							Dir:      "testing_env/modules/t3",
							FileName: "test_3.go",
							Template: "testing_env/templates/usecase_interfaces.text",
							Active:   true,
							Tags: []string{
								"group2",
							},
						},
						{
							Name:     "test 4",
							Dir:      "testing_env/modules/t4",
							FileName: "test_4.go",
							Template: "testing_env/templates/usecase_interfaces.text",
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
				f, err := ioutil.ReadDir("testing_env/modules")
				if assert.Nil(t, err) && assert.Equal(t, 3, len(f)) {
				}
				os.RemoveAll("testing_env/modules")

			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
				global.Tags = []string{"group1"}
			},
		},
		{
			name: "inject code in generate tag",
			args: args{
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					Jobs: []Job{
						{
							Name:       "test",
							Dir:        "testing_env",
							FileName:   "route.go",
							Template:   "testing_env/templates/route.text",
							Active:     true,
							GenerateIn: "route",
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := os.Open("testing_env/route.go")
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
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					Jobs: []Job{
						{
							Name:          "main",
							Dir:           "testing_env/app",
							FileName:      "main.go",
							Template:      "testing_env/templates/main.text",
							Active:        true,
							GenerateIn:    "route",
							SingleExecute: true,
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := os.Open("testing_env/route.go")
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
				yesForAll = utils.BoolP(true)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				GetModelTrait: NewModelFileReader(),
				ReaderMeta:    NewReaderMeta(),
				Decoder:       NewDecoder(tt.args.c),
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
