package company_usecase_v1

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/repository/company_repository"
)

type companyDeleteUsecase struct {
	companyDelete company_repository.CompanyDelete
}

func NewCompanyDeleteUsecase(companyDelete company_repository.CompanyDelete) *companyDeleteUsecase {
	return &companyDeleteUsecase{companyDelete: companyDelete}
}

func (u companyDeleteUsecase) Delete(ctx context.Context, uuid string) (err error) {
	return u.companyDelete.Delete(ctx, uuid)
}
