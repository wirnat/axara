package branch_repository

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/response/branch_response"
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
	Paginate(ctx context.Context, param branch_request.BranchParam) (branch_response.BranchPaginate, error)
}
