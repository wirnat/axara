package company_usecase_v1

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/company/company_repository"
	"github.com/wirnat/aksara-cli/company/company_request"
)

type companyUsecase struct {
	companyStore  company_repository.CompanyStore
	companyUpdate company_repository.CompanyUpdate}

func NewCompanyUsecase(companyStore company_repository.CompanyStore, companyUpdate company_repository.CompanyUpdate) *companyUsecase {
	return &companyUsecase{companyStore: companyStore,companyUpdate: companyUpdate}
}

func (u companyUsecase) Store(ctx context.Context, param company_request.CompanyStore) (r model.Company, err error) {
     r.ID = param.ID
     r.UUID = param.UUID
     r.CreatedAt = param.CreatedAt
     r.UpdatedAt = param.UpdatedAt
     r.DeletedAt = param.DeletedAt
     r.Name = param.Name

	err = u.companyStore.Store(ctx, &r)
	return
}

func (u companyUsecase) Update(ctx context.Context, param company_request.CompanyStore) (r model.Company, err error) {
     r.ID = param.ID
     r.UUID = param.UUID
     r.CreatedAt = param.CreatedAt
     r.UpdatedAt = param.UpdatedAt
     r.DeletedAt = param.DeletedAt
     r.Name = param.Name

	err = u.companyStore.Store(ctx, &r)
	return
}
