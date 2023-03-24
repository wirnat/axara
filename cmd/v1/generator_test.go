package v1

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/ioutil"
	"os"
	"os/exec"
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
					ModuleTraits: []ModuleTrait{
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
					ModuleTraits: []ModuleTrait{
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
	}
	cmd := exec.Command("axara", "set", "--git-key=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")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				GetModelTrait: NewModelFileReader(),
				ReaderMeta:    NewReaderMeta(),
			}
			tt.init()
			err = g.Generate(tt.args.c)
			assert.Equal(t, tt.wantErr, err)
			if tt.fun != nil {
				tt.fun(t)
			}
		})
	}
}

func Test_generator_generateOnce(t *testing.T) {
	type fields struct {
		GetModelTrait FileModelTrait
		ReaderMeta    ReaderMeta
		Puller        Puller
	}
	type args struct {
		c Constructor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		fun     func(testingT assert.TestingT)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "get from remote",
			fields: fields{
				GetModelTrait: NewModelFileReader(),
				ReaderMeta:    NewReaderMeta(),
				Puller:        NewGitPuller(),
			},
			args: args{
				c: Constructor{
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ResultPath: "",
					ModuleName: "~request~",
					ModuleTraits: []ModuleTrait{
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
					Meta: map[string]string{
						"request": "test",
					},
					Traits: []ModuleTrait{
						{
							Name:   "infrastructure",
							Active: true,
							CMD: []string{
								"axara",
								"get",
								"github.com/wirnat/template-aksara-cli-clean-arch",
								"testing_env/infrastructure",
							},
						},
					},
				},
			},
			wantErr: nil,
			fun: func(t assert.TestingT) {
				f, err := ioutil.ReadDir("testing_env/infrastructure")
				t1 := assert.Condition(t, func() (success bool) {
					return len(f) > 0
				})

				t2 := assert.Nil(t, err)
				if t1 && t2 {
					os.RemoveAll("testing_env/infrastructure")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				GetModelTrait: tt.fields.GetModelTrait,
				ReaderMeta:    tt.fields.ReaderMeta,
			}
			err := g.generateOnce(tt.args.c)
			assert.Nil(t, err)
			if tt.fun != nil {
				tt.fun(t)
			}
		})
	}
}
