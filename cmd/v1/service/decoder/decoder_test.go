package decoder

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1"
	"testing"
)

func Test_decoder_Decode(t *testing.T) {
	type fields struct {
		Constructor v1.Constructor
	}
	type args struct {
		code []string
	}

	builder := v1.ModuleBuilder{
		Constructor: v1.Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "testing_env/model",
			ModuleName: "github.com",
			Jobs: []v1.Job{
				{
					Name: "repository", Dir: "testing_env/modules", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"module_target":         "branch",
				"import_infrastructure": "github.com/test",
				"dir_target":            "~module_target~/~import_infrastructure~",
			},
		},
		ModelTrait: &v1.ModelTrait{
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
				Constructor: builder.Constructor,
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
					"~dir_target~",
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
				"branch/github.com/test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := decoder{
				Construct: &tt.fields.Constructor,
			}
			for i, c := range tt.args.code {
				decoded := d.Decode(c, builder.ModelTrait)
				assert.Equal(t, tt.wantEncoded[i], decoded)
			}
		})
	}

}
