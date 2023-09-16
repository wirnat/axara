package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/model"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/repository/company_repository"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/request/company_request"
	"github.com/google/uuid"
)

type companyStoreUsecase struct {
	companyStore company_repository.CompanyStore
}

func NewCompanyStoreUsecase(companyStore company_repository.CompanyStore) *companyStoreUsecase {
	return &companyStoreUsecase{companyStore: companyStore	}
}

func (u companyStoreUsecase) Store(ctx context.Context, param company_request.CompanyStore) (r model.Company, err error) {
        r.ID = param.ID
        r.UUID = uuid.New().String()
        r.createdAt = param.createdAt
        r.createdBy = param.createdBy
        r.updatedAt = param.updatedAt
        r.updatedBy = param.updatedBy
        r.deletedAt = param.deletedAt
        r.deletedBy = param.deletedBy
        r.name = param.name
        r.parent = param.parent
        r.parentUUID = param.parentUUID
        r.subCompany = param.subCompany
        r.type = param.type

    	err = u.companyStore.Store(ctx, &r)
    	return
}
