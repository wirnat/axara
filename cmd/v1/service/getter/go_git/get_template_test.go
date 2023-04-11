package go_git

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetTemplate_Get(t *testing.T) {
	type args struct {
		url string
		dir string
	}
	tests := []struct {
		name string
		args args
		fn   func(t2 *testing.T, err error)
	}{
		{name: "normal repo", args: struct {
			url string
			dir string
		}{url: "https://github.com/zpqrtbnk/test-repo", dir: "clone1"}, fn: func(t2 *testing.T, err error) {
			f, err := ioutil.ReadDir("clone1")
			if assert.Nil(t, err) && assert.LessOrEqual(t, 1, len(f)) {
				os.RemoveAll("clone1")
			}
			assert.Nil(t, err)
		}},
		{name: "specific version", args: struct {
			url string
			dir string
		}{url: "https://github.com/spotify/git-test@v1.0.0", dir: "clone2"}, fn: func(t2 *testing.T, err error) {
			f, err := ioutil.ReadDir("clone2")
			if assert.Nil(t, err) && assert.LessOrEqual(t, 1, len(f)) {
				os.RemoveAll("clone2")
			}
			assert.Nil(t, err)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGetter()
			err := g.Get(tt.args.url, tt.args.dir)
			tt.fn(t, err)
		})
	}
}
