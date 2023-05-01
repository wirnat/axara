package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/request/company_request"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/response/company_response"
)

type companyFetchUsecase struct {
	companyPaginate company_repository.CompanyPaginate
	companyFetch    company_repository.CompanyFetch
}

func NewCompanyFetchUsecase(companyPaginate company_repository.CompanyPaginate, companyFetch company_repository.CompanyFetch) *companyFetchUsecase {
	return &companyFetchUsecase{companyPaginate: companyPaginate, companyFetch: companyFetch}
}
func (f companyFetchUsecase) Paginate(ctx context.Context, param company_request.CompanyParam) (company_response.CompanyPaginate, error) {
	return f.companyPaginate.Paginate(ctx, param)
}

func (f companyFetchUsecase) Fetch(ctx context.Context, param company_request.CompanyParam) ([]model.Company, error) {
	return f.companyFetch.Fetch(ctx, param)
}
