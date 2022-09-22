package branch_usecase

import (
	"github.com/wirnat/aksara-cli/example/model"
	"context"
	"github.com/wirnat/aksara-cli/modules/branch/branch_request"
)

type BranchUsecase interface {
	Store(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
	Update(ctx context.Context, req branch_request.BranchStore) (r model.Branch, err error)
}
