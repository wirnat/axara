package branch_usecase

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/response/branch_response"
)

type BranchStore interface {
	Store(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
}

type BranchUpdate interface {
	Update(ctx context.Context, req branch_request.BranchUpdate) (r model.Branch, err error)
}

type BranchFetch interface {
	Fetch(ctx context.Context, param branch_request.BranchParam) ([]model.Branch, error)
	Paginate(ctx context.Context, param branch_request.BranchParam) (branch_response.BranchPaginate, error)
}

type BranchGet interface {
	Get(ctx context.Context, param branch_request.BranchParam) (model.Branch, error)
}

type BranchDelete interface {
	Delete(ctx context.Context, uuid string) error
}
