package company_usecase

import (
    "context"
    "github.com/wirnat/aksara-cli/example/model"
    "github.com/wirnat/aksara-cli//company/request/company_request"

)

type CompanyUsecase interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
	Update(ctx context.Context, req company_request.CompanyUpdate) (r model.Company, err error)
}
