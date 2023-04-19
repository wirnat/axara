package reader

import (
	"github.com/stretchr/testify/assert"
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/fs"
	"os"
	"testing"
)

func Test_readerFileTs_GetModelTrait(t *testing.T) {
	f, err := os.Stat("test/company.ts")
	if err != nil {
		panic(err)
	}

	type args struct {
		file fs.FileInfo
		c    v1.Constructor
	}
	tests := []struct {
		name string
		args args
		init func(t2 *testing.T)
		res  func(t2 *testing.T, trait *v1.ModelTrait, err error)
	}{
		{
			name: "Read TS Model",
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
					}},
			},
			init: func(t2 *testing.T) {
				global.ExecuteModels = []string{"Company"}
			},
			res: func(t2 *testing.T, trait *v1.ModelTrait, err error) {
				assert.Nil(t, err)
				assert.NotNil(t2, trait.ModelFields[0].Meta)
				assert.Len(t2, trait.ModelFields, 13)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := readerFileTs{}
			tt.init(t)
			gotModelTrait, err := r.GetModelTrait(tt.args.file, tt.args.c)
			tt.res(t, gotModelTrait, err)
		})
	}
}
