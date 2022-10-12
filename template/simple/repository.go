package simple

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
)

type BranchRepository interface {
	BranchRepositoryStore
	BranchRepositoryUpdate
	BranchRepositoryGet
	BranchRepositoryFetch
}

type BranchRepositoryStore interface {
	Store(ctx context.Context, branch *model.Branch) error
}

type BranchRepositoryUpdate interface {
	Update(ctx context.Context, branch *model.Branch) error
}

type BranchRepositoryGet interface {
	Get(ctx context.Context, uuid string) (model.Branch, error)
}

type BranchRepositoryFetch interface {
	Fetch(ctx context.Context) ([]model.Branch, error)
}
