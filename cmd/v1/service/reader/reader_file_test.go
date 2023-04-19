package reader

import (
	"github.com/stretchr/testify/assert"
	"github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"os"
	"testing"
)

func Test_file_GetModelTrait(t *testing.T) {
	f, err := os.Stat("test/company.go")
	if err != nil {
		panic(err)
	}

	type args struct {
		file fs.FileInfo
		c    v1.Constructor
	}
	tests := []struct {
		name           string
		args           args
		wantModelTrait *v1.ModelTrait
		wantErr        assert.ErrorAssertionFunc
		fun            func(t assert.TestingT, trait *v1.ModelTrait, err error)
		init           func()
	}{
		{
			name: "Read file",
			args: args{
				file: f,
				c: v1.Constructor{
					Key:        "ᬅᬓ᭄ᬱᬭ",
					ModelPath:  "test",
					ModuleName: "github.com",
					Jobs: []v1.Job{
						{
							Name: "repository", Dir: "testing_env/modules", FileName: "branch.go",
							Template: "",
						},
					},
				},
			},
			wantModelTrait: nil,
			wantErr:        nil,
			fun: func(t assert.TestingT, trait *v1.ModelTrait, err error) {
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
