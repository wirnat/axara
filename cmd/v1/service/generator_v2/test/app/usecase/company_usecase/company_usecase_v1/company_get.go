package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/company_request"
)

type companyGetUsecase struct {
	companyGet company_repository.CompanyGet
}

func NewCompanyGetUsecase(companyGet company_repository.CompanyGet) *companyGetUsecase {
	return &companyGetUsecase{companyGet: companyGet}
}

func (e companyGetUsecase) Get(ctx context.Context, param company_request.CompanyParam) (model.Company, error) {
	return e.companyGet.Get(ctx, param)
}
