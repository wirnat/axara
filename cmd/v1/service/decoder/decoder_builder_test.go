package decoder

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1"
	"testing"
)

func Test_decoderBuilder_DecodeBuilder(t *testing.T) {
	type fields struct{}
	builder := v1.ModuleBuilder{
		Constructor: v1.Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "branch",
			ModuleName: "github.com",
			Jobs: []v1.Job{
				{
					Name: "repository", Dir: "testing_env/modules/~meta~/~model_snake~/~model~", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"meta": "test",
			},
		},
		ModelTrait: &v1.ModelTrait{
			ModelFields: nil,
			Model:       "Branch",
			ModelSnake:  "branch",
			ModelCamel:  "branch",
		},
	}
	expected := v1.ModuleBuilder{
		Constructor: v1.Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "branch",
			ModuleName: "github.com",
			Jobs: []v1.Job{
				{
					Name: "repository", Dir: "testing_env/modules/test/branch/Branch", FileName: "branch.go",
					Template: "",
				},
			},
			Meta: map[string]string{
				"meta": "test",
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
		name   string
		fields fields
		wantR  v1.ModuleBuilder
	}{
		{
			name:   "Embed must valid depend on the Decoder",
			fields: fields{},
			wantR:  expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewDecoder(&builder.Constructor)
			result := b.DecodeBuilder(builder)
			assert.Equalf(t, tt.wantR, result, "DecodeBuilder()")
		})
	}
}
