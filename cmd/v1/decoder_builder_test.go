package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decoderBuilder_DecodeBuilder(t *testing.T) {
	type fields struct {
		Decoder Decoder
	}
	builder := ModuleBuilder{
		Constructor: Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "branch",
			ModuleName: "github.com",
			ModuleTraits: []ModuleTrait{
				{
					Name: "repository", Dir: "testing_env/modules/~meta~/~model_snake~/~model~", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"meta": "test",
			},
		},
		ModelTrait: ModelTrait{
			ModelFields: nil,
			Model:       "Branch",
			ModelSnake:  "branch",
			ModelCamel:  "branch",
		},
	}
	expected := ModuleBuilder{
		Constructor: Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "branch",
			ModuleName: "github.com",
			ModuleTraits: []ModuleTrait{
				{
					Name: "repository", Dir: "testing_env/modules/test/branch/Branch", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"meta": "test",
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
		name   string
		fields fields
		wantR  ModuleBuilder
	}{
		{
			name: "Embed must valid depend on the Decoder",
			fields: fields{
				Decoder: NewDecoder(builder.Constructor),
			},
			wantR: expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := decoderBuilder{
				Decoder: tt.fields.Decoder,
			}
			result := b.DecodeBuilder(builder)
			assert.Equalf(t, tt.wantR, result, "DecodeBuilder()")
		})
	}
}
