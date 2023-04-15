package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/branch_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
	"github.com/google/uuid"
)

type branchStoreUsecase struct {
	branchStore branch_repository.BranchStore
}

func NewBranchStoreUsecase(branchStore branch_repository.BranchStore) *branchStoreUsecase {
	return &branchStoreUsecase{branchStore: branchStore	}
}

func (u branchStoreUsecase) Store(ctx context.Context, param branch_request.BranchStore) (r model.Branch, err error) {
        r.ID = param.ID
        r.UUID = uuid.New().String()
        r.CreatedAt = param.CreatedAt
        r.UpdatedAt = param.UpdatedAt
        r.DeletedAt = param.DeletedAt
        r.CompanyID = param.CompanyID
        r.Name = param.Name
        r.Description = param.Description

    	err = u.branchStore.Store(ctx, &r)
    	return
}
