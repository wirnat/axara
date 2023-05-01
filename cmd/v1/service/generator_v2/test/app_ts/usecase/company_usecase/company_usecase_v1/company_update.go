package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/request/company_request"
)

type companyUpdateUsecase struct {
	companyUpdate company_repository.CompanyUpdate
	companyGet    company_repository.CompanyGet
}

func NewCompanyUpdateUsecase(companyUpdate company_repository.CompanyUpdate, companyGet company_repository.CompanyGet) *companyUpdateUsecase {
	return &companyUpdateUsecase{companyUpdate: companyUpdate,companyGet: companyGet}
}

func (u companyUpdateUsecase) Update(ctx context.Context, param company_request.CompanyUpdate) (r model.Company, err error) {
    	     r.ID = param.ID
    	     r.createdAt = param.createdAt
    	     r.createdBy = param.createdBy
    	     r.updatedAt = param.updatedAt
    	     r.updatedBy = param.updatedBy
    	     r.deletedAt = param.deletedAt
    	     r.deletedBy = param.deletedBy
    	     r.name = param.name
    	     r.parent = param.parent
    	     r.parentUUID = param.parentUUID
    	     r.subCompany = param.subCompany
    	     r.type = param.type

    	err = u.companyUpdate.Update(ctx, &r,company_request.CompanyParam{UUID: &param.UUID})
    	if err!=nil{
    	    return
    	}

    	r, err = u.companyGet.Get(ctx, company_request.CompanyParam{
        	UUID: &param.UUID,
        })

    	return
}
