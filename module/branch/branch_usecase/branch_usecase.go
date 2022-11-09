package branch_usecase

import (
    "context"
    "github.com/wirnat/aksara-cli/example/model"
    "github.com/wirnat/aksara-cli/module/branch/request/branch_request"
)

type BranchUsecase interface {
	Store(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
	Update(ctx context.Context, req branch_request.BranchUpdate) (r model.Branch, err error)
}
