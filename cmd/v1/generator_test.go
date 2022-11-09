package v1

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1/errors"
	"io/ioutil"
	"os"
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
		fun     func(t assert.TestingT)
	}{
		{
			name: "Key not equal with ᬅᬓ᭄ᬱᬭ",
			args: args{
				c: Constructor{
					Key:           "",
					ModelPath:     "",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: nil,
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.InvalidKey,
		},
		{
			name: "Construct model traits<1 return error: nothing todo ",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: nil,
				},
			},
			wantErr: errors.NothingTodo,
		},
		{
			name: "Get directory from model, if not valid return error: no model found ",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model_empty",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: []string{"Company"},
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.NoModelFound,
		},
		{
			name: "No model contain ~model~ inside the directory, return no model contain ~model~",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model_without_flag",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: []string{"Company"},
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.NoModelCanExecute,
		},
		{
			name: "~model~ was found, but not has any value, return error: InvalidModelFlag",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model_with_invalid_value",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: []string{"Company"},
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.InvalidModelFlag,
		},
		{
			name: "~model~ was found, but not has ~end_model~ comment",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model_without_end",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: []string{"Company"},
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.NoEndModelFound,
		},
		{
			name: "No match execute model",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model",
					ResultPath:    "",
					ModuleName:    "",
					ExecuteModels: []string{"Company"},
					ModuleTraits: []ModuleTrait{
						{
							Name:     "test",
							Dir:      "",
							FileName: "test",
							Template: "test",
						},
					},
				},
			},
			wantErr: errors.NoModelCanExecute,
		},
		{
			name: "total generated directory must equal with execute trait on the model",
			args: args{
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model",
					ResultPath:    "",
					ModuleName:    "~request~",
					ExecuteModels: []string{"Branch"},
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
		},
		{
			name: "only generate active trait",
			args: args{
				c: Constructor{
					Meta: map[string]string{
						"request": "test",
					},
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model",
					ResultPath:    "",
					ModuleName:    "~request~",
					ExecuteModels: []string{"Branch"},
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
		},
		{
			name: "generate trait with remote",
			args: args{
				c: Constructor{
					Key:           "ᬅᬓ᭄ᬱᬭ",
					ModelPath:     "testing_env/model",
					ResultPath:    "",
					ModuleName:    "~request~",
					ExecuteModels: []string{"Branch"},
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
							Dir:    "testing_env/infrastructure",
							Remote: "github.com/wirnat/template-aksara-cli-clean-arch",
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
				GetModelTrait: NewModelFileReader(),
				ReaderMeta:    NewReaderMeta(),
				Puller:        NewGitPuller(),
			}
			err := g.Generate(tt.args.c)
			assert.Equal(t, tt.wantErr, err)
			if tt.fun != nil {
				tt.fun(t)
			}
		})
	}
}
