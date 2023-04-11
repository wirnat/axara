package branch_usecase

import (
    "context"
    "test 1/../../spam/testing_env/model"
    "test 1/../../spam/testing_env/modules/branch/request/branch_request"

)

type BranchUsecase interface {
	Store(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
	Update(ctx context.Context, req branch_request.BranchUpdate) (r model.Branch, err error)
}
