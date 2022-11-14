package company_usecase

import (
    "context"
    "github.com/wirnat/axara/example/model"
    "github.com/wirnat/axara/module/company/request/company_request"
)

type CompanyUsecase interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
	Update(ctx context.Context, req company_request.CompanyUpdate) (r model.Company, err error)
}
