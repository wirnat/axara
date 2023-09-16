package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/branch_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
)

type branchUpdateUsecase struct {
	branchUpdate branch_repository.BranchUpdate
	branchGet    branch_repository.BranchGet
}

func NewBranchUpdateUsecase(branchUpdate branch_repository.BranchUpdate, branchGet branch_repository.BranchGet) *branchUpdateUsecase {
	return &branchUpdateUsecase{branchUpdate: branchUpdate,branchGet: branchGet}
}

func (u branchUpdateUsecase) Update(ctx context.Context, param branch_request.BranchUpdate) (r model.Branch, err error) {

    	err = u.branchUpdate.Update(ctx, &r,branch_request.BranchParam{UUID: &param.UUID})
    	if err!=nil{
    	    return
    	}

    	r, err = u.branchGet.Get(ctx, branch_request.BranchParam{
        	UUID: &param.UUID,
        })

    	return
}
