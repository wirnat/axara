package company_usecase

import (
    "gitlab.com/wirawirw/aksara-cli/example/model"
    "context"
    "gitlab.com/wirawirw/aksara-cli/modules/company/company_request"
)

type CompanyUsecase interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
	Update(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
}
