package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/company_request"
)

type companyUpdateUsecase struct {
	companyUpdate company_repository.CompanyUpdate
	companyGet    company_repository.CompanyGet
}

func NewCompanyUpdateUsecase(companyUpdate company_repository.CompanyUpdate, companyGet company_repository.CompanyGet) *companyUpdateUsecase {
	return &companyUpdateUsecase{companyUpdate: companyUpdate,companyGet: companyGet}
}

func (u companyUpdateUsecase) Update(ctx context.Context, param company_request.CompanyUpdate) (r model.Company, err error) {

    	err = u.companyUpdate.Update(ctx, &r,company_request.CompanyParam{UUID: &param.UUID})
    	if err!=nil{
    	    return
    	}

    	r, err = u.companyGet.Get(ctx, company_request.CompanyParam{
        	UUID: &param.UUID,
        })

    	return
}
