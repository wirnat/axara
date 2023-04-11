package v1

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"os"
	"testing"
)

func Test_file_GetModelTrait(t *testing.T) {
	f, err := os.Stat("testing_env/model/company.go")
	if err != nil {
		panic(err)
	}

	type args struct {
		file fs.FileInfo
		c    Constructor
	}
	tests := []struct {
		name           string
		args           args
		wantModelTrait *ModelTrait
		wantErr        assert.ErrorAssertionFunc
		fun            func(t assert.TestingT, trait *ModelTrait, err error)
		init           func()
	}{
		{
			name: "Read file",
			args: args{
				file: f,
				c: Constructor{
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "testing_env/model",
					ModuleName: "github.com",
					Jobs: []Job{
						{
							Name: "repository", Dir: "testing_env/modules", FileName: "branch.go",
							Template: "",
						},
					},
				},
			},
			wantModelTrait: nil,
			wantErr:        nil,
			fun: func(t assert.TestingT, trait *ModelTrait, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, trait)
				assert.Len(t, trait.ModelFields, 6)
			},
			init: func() {
				global.ExecuteModels = []string{"Company"}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := file{}
			tt.init()
			gotModelTrait, err := g.GetModelTrait(tt.args.file, tt.args.c)
			tt.fun(t, gotModelTrait, err)
		})
	}
}
