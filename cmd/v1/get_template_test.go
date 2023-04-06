package v1

import (
	"github.com/stretchr/testify/assert"
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
	}{
		{name: "normal repo", args: struct {
			url string
			dir string
		}{url: "https://github.com/zpqrtbnk/test-repo", dir: "clone"}},
		{name: "specific version", args: struct {
			url string
			dir string
		}{url: "https://github.com/spotify/git-test@v1.0.0", dir: "clone/v1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGetter()
			err := g.Get(tt.args.url, tt.args.dir)
			assert.Nil(t, err)
		})
	}
}
