package company_usecase_v1

import (
	"context"
	"github.com/google/uuid"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/module_fix/company/company_usecase"
	"github.com/wirnat/aksara-cli/module_fix/company/repository/company_repository"
	"github.com/wirnat/aksara-cli/module_fix/company/request/company_request"
)

type companyUsecase struct {
	companyStore  company_repository.CompanyStore
	companyGet    company_repository.CompanyGet
	companyUpdate company_repository.CompanyUpdate
}

func NewCompanyUsecase(
	companyStore company_repository.CompanyStore,
	companyUpdate company_repository.CompanyUpdate,
	companyGet company_repository.CompanyGet,
) company_usecase.CompanyUsecase {
	return &companyUsecase{companyStore: companyStore, companyUpdate: companyUpdate, companyGet: companyGet}
}

func (u companyUsecase) Store(ctx context.Context, Param company_request.CompanyStore) (r model.Company, err error) {
	r.ID = Param.ID
	r.UUID = uuid.New().String()
	r.Name = Param.Name

	err = u.companyStore.Store(ctx, &r)
	return
}

func (u companyUsecase) Update(ctx context.Context, Param company_request.CompanyUpdate) (r model.Company, err error) {
	r.ID = Param.ID
	r.Name = Param.Name

	err = u.companyUpdate.Update(ctx, &r, company_request.CompanyParam{UUID: &r.UUID})
	if err != nil {
		return
	}

	r, err = u.companyGet.Get(ctx, company_request.CompanyParam{
		UUID: &r.UUID,
	})

	return
}
