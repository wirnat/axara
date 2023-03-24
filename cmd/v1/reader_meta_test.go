package v1

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func Test_readerMeta_GetMeta(t *testing.T) {
	f, err := os.Stat("testing_env/model/branch.go")
	if err != nil {
		panic(err)
	}

	builder := ModuleBuilder{
		Constructor: Constructor{
			Key:        "ᬅᬓ᭄ᬱᬭ",
			ModelPath:  "testing_env/model",
			ModuleName: "github.com",
			ModuleTraits: []ModuleTrait{
				{
					Name: "repository", Dir: "testing_env/modules", FileName: "branch.go",
					Template: "",
				},
			},
			Models: map[string]map[string]string{
				"Branch": {
					"module": "branch",
				},
			},
		},
		ModelTrait: ModelTrait{
			ModelFields: nil,
			Model:       "Branch",
			ModelSnake:  "branch",
			ModelCamel:  "branch",
		},
	}

	type args struct {
		file fs.FileInfo
		c    Constructor
	}
	tests := []struct {
		name     string
		args     args
		wantMeta map[string]string
		wantErr  error
	}{
		{
			name: "Get meta",
			args: args{
				file: f,
				c:    builder.Constructor,
			},
			wantMeta: map[string]string{"module": "branch"},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := readerMeta{}
			gotMeta, err := m.GetMeta(tt.args.file, tt.args.c, "Branch")
			assert.Equal(t, tt.wantErr, err)
			assert.Equalf(t, tt.wantMeta, gotMeta, "GetMeta()")
		})
	}
}
