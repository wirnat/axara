package company_repository

import (
	"context"
	"gitlab.com/wirawirw/aksara-cli/example/model"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_request"
)

type Company interface {
	CompanyFetch
	CompanyGet
	CompanyStore
	CompanyUpdate
	CompanyDelete
}

type CompanyFetch interface {
	Fetch(ctx context.Context, param company_request.CompanyParam) ([]model.Company, error)
}

type CompanyGet interface {
	Get(ctx context.Context, param company_request.CompanyParam) (model.Company, error)
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
