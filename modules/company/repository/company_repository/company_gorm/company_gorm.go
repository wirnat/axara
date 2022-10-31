package company_gorm

import (
	"context"
	"fmt"
	"github.com/wirnat/aksara-cli/example/model"
	"github.com/wirnat/aksara-cli/modules/company/request/company_request"
	"github.com/wirnat/aksara-cli/modules/company/response"
	"github.com/wirnat/aksara-cli/template/clean_architecture/infrastructure/paginator"
	"gorm.io/gorm"
)

type companyGorm struct {
	db gorm.DB
}

func New(db gorm.DB) *companyGorm {
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
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return c.db.Create(&company).Error
}

func (c companyGorm) Update(ctx context.Context, company *model.Company, condition ...company_request.CompanyParam) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	for _, p := range condition {
        filter(&c.db, p)
    }

	return c.db.Updates(&company).Error
}

func (c companyGorm) Delete(ctx context.Context, uuid string) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Delete(model.Company{}, fmt.Sprintf("uuid=%v", uuid)).Error
}

func (c companyGorm) Paginate(ctx context.Context, param company_request.CompanyParam) (response.CompanyPaginate, error) {
	res := response.CompanyPaginate{
		Pagination: &paginator.Pagination{
			PaginateReq: param.PaginateReq,
		},
	}

	filter(&c.db, param)
	err := c.db.Debug().Scopes(paginator.PaginateV2(&res.Result, &c.db, res.Pagination)).Find(&res.Result).Error
	if err != nil {
		return response.CompanyPaginate{}, err
	}

	return res, nil
}

