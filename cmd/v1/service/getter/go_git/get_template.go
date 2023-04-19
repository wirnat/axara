package go_git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"os"
	"strings"
)

type Getter interface {
	Get(url string, dir string) error
}

type goGit struct{}

func NewGetter() *goGit {
	return &goGit{}
}

func (g goGit) Get(url string, dir string) error {
	hV := strings.Split(url, "@")
	var rf plumbing.ReferenceName
	if len(hV) == 2 {
		rf = plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", hV[1]))
		url = hV[0]
	}

	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_, err = git.PlainClone(dir, false, &git.CloneOptions{
			URL:           url,
			Progress:      os.Stdout,
			Depth:         0, // download the entire history
			ReferenceName: rf,
		})
	} else {
		// open repository if directory already exists
		repo, err := git.PlainOpen(dir)
		if err != nil {
			panic(err)
		}
		// get local reference
		w, err := repo.Worktree()
		if err != nil {
			panic(err)
		}

		// pull changes from remote
		err = w.Pull(&git.PullOptions{
			RemoteName: "origin",
			Progress:   os.Stdout,
		})
		if err != nil {
			if err != git.NoErrAlreadyUpToDate {
				panic(err)
			}
		}
	}

	if err != nil {
		return err
	}

	return nil
}
