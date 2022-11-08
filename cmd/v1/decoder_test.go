package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decoder_Decode(t *testing.T) {
	type fields struct {
		Builder ModuleBuilder
	}
	type args struct {
		code []string
	}

	builder := ModuleBuilder{
		Constructor: Constructor{
			Key:           "ᬅᬓ᭄ᬱᬭ",
			ModelPath:     "testing_env/model",
			ModuleName:    "github.com",
			ExecuteModels: []string{"Branch"},
			ModuleTraits: []ModuleTrait{
				{
					Name: "repository", Dir: "testing_env/modules", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"module_target":         "branch",
				"import_infrastructure": "github.com/test",
			},
		},
		ModelTrait: ModelTrait{
			ModelFields: nil,
			Model:       "Branch",
			ModelSnake:  "branch",
			ModelCamel:  "branch",
		},
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		wantEncoded []string
	}{
		{
			name: "All decoded code must match with builder",
			fields: fields{
				Builder: builder,
			},
			args: args{
				code: []string{
					"~model_path~",
					"~model_snake~",
					"~model~",
					"~model_camel~",
					"~module_name~",
					"~model_path~",
					"~module_target~",
					"~import_infrastructure~",
				},
			},
			wantEncoded: []string{
				builder.ModelPath,
				builder.ModelSnake,
				builder.Model,
				builder.ModelCamel,
				builder.ModuleName,
				builder.ModelPath,
				"branch",
				"github.com/test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := decoder{
				Builder: tt.fields.Builder,
			}
			for i, c := range tt.args.code {
				decoded := d.Decode(c)
				assert.Equal(t, tt.wantEncoded[i], decoded)
			}
		})
	}

}
