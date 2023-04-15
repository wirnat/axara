package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/branch_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
)

type branchGetUsecase struct {
	branchGet branch_repository.BranchGet
}

func NewBranchGetUsecase(branchGet branch_repository.BranchGet) *branchGetUsecase {
	return &branchGetUsecase{branchGet: branchGet}
}

func (e branchGetUsecase) Get(ctx context.Context, param branch_request.BranchParam) (model.Branch, error) {
	return e.branchGet.Get(ctx, param)
}
