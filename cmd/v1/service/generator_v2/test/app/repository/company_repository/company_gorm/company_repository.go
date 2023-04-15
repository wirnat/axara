package company_gorm

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/model"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/request/company_request"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/domain/response/company_response"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/paginator"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/db_transaction/gorm_transaction"
	"gorm.io/gorm"
)

type companyGorm struct {
	db gorm.DB
}

func New(db gorm.DB) *companyGorm {
    db.AutoMigrate(&model.Company{})
	return &companyGorm{db: db}
}

func (c companyGorm) Fetch(ctx context.Context, Param company_request.CompanyParam) (res []model.Company, err error) {
	filter(&c.db, Param)
	err = c.db.Find(&res).Error
	return
}

func (c companyGorm) Get(ctx context.Context, Param company_request.CompanyParam) (res model.Company, err error) {
	filter(&c.db, Param)
	err = c.db.First(&res).Error
	return
}

func (c companyGorm) Store(ctx context.Context, company *model.Company) error {
	db, err :=  gorm_transaction.GetTx(ctx)
	if err == nil {
		c.db = *db
	}

	return c.db.Create(&company).Error
}

func (c companyGorm) Update(ctx context.Context, company *model.Company, condition ...company_request.CompanyParam) error {
	db, err :=  gorm_transaction.GetTx(ctx)
	if err == nil {
		c.db = *db
	}

	for _, p := range condition {
        filter(&c.db, p)
    }

	return c.db.Updates(&company).Error
}

func (c companyGorm) Delete(ctx context.Context, uuid string) error {
	db, err :=  gorm_transaction.GetTx(ctx)
	if err == nil {
		c.db = *db
	}
	return c.db.Debug().Where("uuid=?", uuid).Delete(&model.Company{}).Error
}

func (c companyGorm) Paginate(ctx context.Context, param company_request.CompanyParam) (company_response.CompanyPaginate, error) {
	res := company_response.CompanyPaginate{
		Pagination: &paginator.Pagination{
			PaginateReq: param.PaginateReq,
		},
	}

	filter(&c.db, param)
	err := c.db.Debug().Scopes(paginator.PaginateV2(&res.Result, &c.db, res.Pagination)).Find(&res.Result).Error
	if err != nil {
		return company_response.CompanyPaginate{}, err
	}

	return res, nil
}

