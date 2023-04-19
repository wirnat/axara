package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/branch_repository"
)

type branchDeleteUsecase struct {
	branchDelete branch_repository.BranchDelete
}

func NewBranchDeleteUsecase(branchDelete branch_repository.BranchDelete) *branchDeleteUsecase {
	return &branchDeleteUsecase{branchDelete: branchDelete}
}

func (u branchDeleteUsecase) Delete(ctx context.Context, uuid string) (err error) {
	return u.branchDelete.Delete(ctx, uuid)
}
