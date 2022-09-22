package branch_repository

import (
	"context"
	"gitlab.com/wirawirw/aksara-cli/example/model"
	"gitlab.com/wirawirw/aksara-cli/modules/branch/branch_request"
)

type Branch interface {
	BranchFetch
	BranchGet
	BranchStore
	BranchUpdate
	BranchDelete
}

type BranchFetch interface {
	Fetch(ctx context.Context, param branch_request.BranchParam) ([]model.Branch, error)
}

type BranchGet interface {
	Get(ctx context.Context, param branch_request.BranchParam) (model.Branch, error)
}

type BranchStore interface {
	Store(ctx context.Context, Branch *model.Branch) error
}

type BranchUpdate interface {
	Update(ctx context.Context, Branch *model.Branch, condition ...branch_request.BranchParam) error
}

type BranchDelete interface {
	Delete(ctx context.Context, uuid string) error
}
