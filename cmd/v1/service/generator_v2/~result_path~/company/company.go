package company_usecase

import (
    "context"
    "test 1/../../spam/testing_env/model"
    "test 1/../../spam/testing_env/modules/company/request/company_request"

)

type CompanyUsecase interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
	Update(ctx context.Context, req company_request.CompanyUpdate) (r model.Company, err error)
}
