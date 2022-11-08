package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/google/uuid"
   "github.com/wirnat/aksara-cli/module/branch/branch_usecase"
   "github.com/wirnat/aksara-cli/module/branch/repository/branch_repository"
   "github.com/wirnat/aksara-cli/module/branch/request/branch_request"
)

type branchUsecase struct {
	branchStore  branch_repository.BranchStore
	branchGet    branch_repository.BranchGet
	branchUpdate branch_repository.BranchUpdate}

func NewBranchUsecase(
    branchStore branch_repository.BranchStore,
    branchUpdate branch_repository.BranchUpdate,
    branchGet    branch_repository.BranchGet,
    ) branch_usecase.BranchUsecase {
	return &branchUsecase{branchStore: branchStore,branchUpdate: branchUpdate,branchGet: branchGet,
	}
}

func (u branchUsecase) Store(ctx context.Context, Param branch_request.BranchStore) (r model.Branch, err error) {
     r.ID = Param.ID
	 r.UUID = uuid.New().String()
     r.CompanyID = Param.CompanyID
     r.Name = Param.Name
     r.Description = Param.Description

	err = u.branchStore.Store(ctx, &r)
	return
}

func (u branchUsecase) Update(ctx context.Context, Param branch_request.BranchUpdate) (r model.Branch, err error) {
	     r.ID = Param.ID
	     r.CompanyID = Param.CompanyID
	     r.Name = Param.Name
	    r.Description = *Param.Description

	err = u.branchUpdate.Update(ctx, &r,branch_request.BranchParam{UUID: &r.UUID})
	if err!=nil{
	    return
	}

	r, err = u.branchGet.Get(ctx, branch_request.BranchParam{
    	UUID: &r.UUID,
    })

	return
}
