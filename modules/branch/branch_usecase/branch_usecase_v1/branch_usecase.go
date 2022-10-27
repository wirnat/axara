package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/branch/branch_repository"
	"github.com/wirnat/aksara-cli/branch/branch_request"
)

type branchUsecase struct {
	branchStore  branch_repository.BranchStore
	branchUpdate branch_repository.BranchUpdate}

func NewBranchUsecase(branchStore branch_repository.BranchStore, branchUpdate branch_repository.BranchUpdate) *branchUsecase {
	return &branchUsecase{branchStore: branchStore,branchUpdate: branchUpdate}
}

func (u branchUsecase) Store(ctx context.Context, param branch_request.BranchStore) (r model.Branch, err error) {
     r.ID = param.ID
     r.UUID = param.UUID
     r.CreatedAt = param.CreatedAt
     r.UpdatedAt = param.UpdatedAt
     r.DeletedAt = param.DeletedAt
     r.CompanyID = param.CompanyID
     r.Name = param.Name
     r.Description = param.Description

	err = u.branchStore.Store(ctx, &r)
	return
}

func (u branchUsecase) Update(ctx context.Context, param branch_request.BranchStore) (r model.Branch, err error) {
     r.ID = param.ID
     r.UUID = param.UUID
     r.CreatedAt = param.CreatedAt
     r.UpdatedAt = param.UpdatedAt
     r.DeletedAt = param.DeletedAt
     r.CompanyID = param.CompanyID
     r.Name = param.Name
     r.Description = param.Description

	err = u.branchStore.Store(ctx, &r)
	return
}
