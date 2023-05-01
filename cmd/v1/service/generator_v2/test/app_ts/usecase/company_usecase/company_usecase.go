package company_usecase

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/model"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/request/company_request"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/response/company_response"
)

type CompanyStore interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
}

type CompanyUpdate interface {
	Update(ctx context.Context, req company_request.CompanyUpdate) (r model.Company, err error)
}

type CompanyFetch interface {
	Fetch(ctx context.Context, param company_request.CompanyParam) ([]model.Company, error)
	Paginate(ctx context.Context, param company_request.CompanyParam) (company_response.CompanyPaginate, error)
}

type CompanyGet interface {
	Get(ctx context.Context, param company_request.CompanyParam) (model.Company, error)
}

type CompanyDelete interface {
	Delete(ctx context.Context, uuid string) error
}
