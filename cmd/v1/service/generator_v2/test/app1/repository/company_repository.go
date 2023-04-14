package repository

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/model"
	"<no value>"
	"<no value>"
)

type Company interface {
	CompanyFetch
	CompanyGet
	CompanyStore
	CompanyUpdate
	CompanyDelete
	CompanyPaginate
}

type CompanyFetch interface {
	Fetch(ctx context.Context, Param company_request.CompanyParam) ([]model.Company, error)
}

type CompanyGet interface {
	Get(ctx context.Context, Param company_request.CompanyParam) (model.Company, error)
}

type CompanyStore interface {
	Store(ctx context.Context, Company *model.Company) error
}

type CompanyUpdate interface {
	Update(ctx context.Context, Company *model.Company, condition ...company_request.CompanyParam) error
}

type CompanyDelete interface {
	Delete(ctx context.Context, uuid string) error
}

type CompanyPaginate interface {
	Paginate(ctx context.Context, param company_request.CompanyParam) (response.CompanyPaginate, error)
}
