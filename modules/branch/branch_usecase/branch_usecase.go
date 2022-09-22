package branch_usecase

import (
    "gitlab.com/wirawirw/aksara-cli/example/model"
    "context"
    "gitlab.com/wirawirw/aksara-cli/modules/branch/branch_request"
)

type BranchUsecase interface {
	Store(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
	Update(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
}
