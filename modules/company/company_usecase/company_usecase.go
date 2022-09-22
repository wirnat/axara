package company_usecase

import (
	"github.com/wirnat/aksara-cli/example/model"
	"context"
	"github.com/wirnat/aksara-cli/modules/company/company_request"
)

type CompanyUsecase interface {
	Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
	Update(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error)
}
