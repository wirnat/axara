package branch_mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wirnat/axara/example/model"
	"github.com/wirnat/axara/module/branch/request/branch_request"
	"github.com/wirnat/axara/module/branch/response"
)

type BranchMock struct {
	mock.Mock
}

func (o BranchMock) Paginate(ctx context.Context, param branch_request.BranchParam) (res response.BranchPaginate, err error) {
	args := o.Called(ctx, param)
	res, _ = args.Get(0).(response.BranchPaginate)
	err, _ = args.Get(1).(error)
	return
}

func (o BranchMock) Fetch(ctx context.Context, Param branch_request.BranchParam) (res []model.Branch, err error) {
	args := o.Called(ctx, Param)
	res, _ = args.Get(0).([]model.Branch)
	err, _ = args.Get(1).(error)
	return
}

func (o BranchMock) Get(ctx context.Context, Param branch_request.BranchParam) (res model.Branch, err error) {
	args := o.Called(ctx, Param)
	res, _ = args.Get(0).(model.Branch)
	err, _ = args.Get(1).(error)
	return
}

func (o BranchMock) Store(ctx context.Context, branch *model.Branch) (err error) {
	args := o.Called(ctx, branch)
	err, _ = args.Get(0).(error)
	return
}

func (o BranchMock) Update(ctx context.Context, branch *model.Branch, condition ...branch_request.BranchParam) (err error) {
	args := o.Called(ctx, branch, condition)
	err, _ = args.Get(0).(error)
	return
}

func (o BranchMock) Delete(ctx context.Context, uuid string) (err error) {
	args := o.Called(ctx, ctx, uuid)
	err, _ = args.Get(0).(error)
	return
}
