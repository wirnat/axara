package v1

import "github.com/stretchr/testify/mock"

type GitPullerMock struct {
	mock.Mock
}

func (g GitPullerMock) Pull(p string, targetDir string) (err error) {
	args := g.Called(p, targetDir)
	err, _ = args.Get(0).(error)
	return
}
