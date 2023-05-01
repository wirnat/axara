package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/company_request"
	"github.com/google/uuid"
)

type companyStoreUsecase struct {
	companyStore company_repository.CompanyStore
}

func NewCompanyStoreUsecase(companyStore company_repository.CompanyStore) *companyStoreUsecase {
	return &companyStoreUsecase{companyStore: companyStore	}
}

func (u companyStoreUsecase) Store(ctx context.Context, param company_request.CompanyStore) (r model.Company, err error) {

    	err = u.companyStore.Store(ctx, &r)
    	return
}
