package v1

import (
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_gitPuller_Pull(t *testing.T) {
	type fields struct {
		client *github.Client
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				path: "",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetDir := "testing_env/storage"
			g := NewGitPuller()
			err := g.Pull("github.com/rusli-nasir/gcp-kubernetes-bucket-nfs-server", targetDir)
			f, err := ioutil.ReadDir(targetDir)
			t1 := assert.Nil(t, err)
			t2 := assert.Condition(t, func() (success bool) {
				if len(f) > 0 {
					return true
				}
				return false
			})
			if t1 && t2 {
				os.RemoveAll(targetDir)
			}
		})
	}
}
