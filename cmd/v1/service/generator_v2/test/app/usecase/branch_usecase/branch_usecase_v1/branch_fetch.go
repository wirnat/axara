package branch_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/branch_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/branch_request"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/response/branch_response"
)

type branchFetchUsecase struct {
	branchPaginate branch_repository.BranchPaginate
	branchFetch    branch_repository.BranchFetch
}

func NewBranchFetchUsecase(branchPaginate branch_repository.BranchPaginate, branchFetch branch_repository.BranchFetch) *branchFetchUsecase {
	return &branchFetchUsecase{branchPaginate: branchPaginate, branchFetch: branchFetch}
}
func (f branchFetchUsecase) Paginate(ctx context.Context, param branch_request.BranchParam) (branch_response.BranchPaginate, error) {
	return f.branchPaginate.Paginate(ctx, param)
}

func (f branchFetchUsecase) Fetch(ctx context.Context, param branch_request.BranchParam) ([]model.Branch, error) {
	return f.branchFetch.Fetch(ctx, param)
}
