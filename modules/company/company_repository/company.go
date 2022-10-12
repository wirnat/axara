package company_repository

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
)

type CompanyRepository interface {
	CompanyRepositoryStore
	CompanyRepositoryUpdate
	CompanyRepositoryGet
	CompanyRepositoryFetch
}

type CompanyRepositoryStore interface {
	Store(ctx context.Context, branch *model.Company) error
}

type CompanyRepositoryUpdate interface {
	Update(ctx context.Context, branch *model.Company) error
}

type CompanyRepositoryGet interface {
	Get(ctx context.Context, uuid string) (model.Company, error)
}

type CompanyRepositoryFetch interface {
	Fetch(ctx context.Context) ([]model.Company, error)
}
