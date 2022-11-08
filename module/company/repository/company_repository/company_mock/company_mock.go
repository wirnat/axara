package company_mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/module/company/request/company_request"
	"github.com/wirnat/aksara-cli/module/company/response"
)

type CompanyMock struct {
	mock.Mock
}

func (o CompanyMock) Paginate(ctx context.Context, param company_request.CompanyParam) (res response.CompanyPaginate, err error) {
	args := o.Called(ctx, param)
	res, _ = args.Get(0).(response.CompanyPaginate)
	err, _ = args.Get(1).(error)
	return
}

func (o CompanyMock) Fetch(ctx context.Context, Param company_request.CompanyParam) (res []model.Company, err error) {
	args := o.Called(ctx, Param)
	res, _ = args.Get(0).([]model.Company)
	err, _ = args.Get(1).(error)
	return
}

func (o CompanyMock) Get(ctx context.Context, Param company_request.CompanyParam) (res model.Company, err error) {
	args := o.Called(ctx, Param)
	res, _ = args.Get(0).(model.Company)
	err, _ = args.Get(1).(error)
	return
}

func (o CompanyMock) Store(ctx context.Context, company *model.Company) (err error) {
	args := o.Called(ctx, company)
	err, _ = args.Get(0).(error)
	return
}

func (o CompanyMock) Update(ctx context.Context, company *model.Company, condition ...company_request.CompanyParam) (err error) {
	args := o.Called(ctx, company, condition)
	err, _ = args.Get(0).(error)
	return
}

func (o CompanyMock) Delete(ctx context.Context, uuid string) (err error) {
	args := o.Called(ctx, ctx, uuid)
	err, _ = args.Get(0).(error)
	return
}
