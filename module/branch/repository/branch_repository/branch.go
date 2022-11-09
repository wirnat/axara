package branch_repository

import (
	"context"
	"github.com/wirnat/axara/example/model"
	"github.com/wirnat/axara/module/branch/request/branch_request"
	"github.com/wirnat/axara/module/branch/response"
)

type Branch interface {
	BranchFetch
	BranchGet
	BranchStore
	BranchUpdate
	BranchDelete
	BranchPaginate
}

type BranchFetch interface {
	Fetch(ctx context.Context, Param branch_request.BranchParam) ([]model.Branch, error)
}

type BranchGet interface {
	Get(ctx context.Context, Param branch_request.BranchParam) (model.Branch, error)
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

type BranchPaginate interface {
	Paginate(ctx context.Context, param branch_request.BranchParam) (response.BranchPaginate, error)
}
