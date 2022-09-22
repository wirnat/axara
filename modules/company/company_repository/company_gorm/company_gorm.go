package company_gorm

import (
	"context"
	"fmt"
	"gitlab.com/wirawirw/aksara-cli/example/model"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_request"
	"gorm.io/gorm"
)

type companyGorm struct {
	db gorm.DB
}

func New(db gorm.DB) *companyGorm {
	return &companyGorm{db: db}
}

func (c companyGorm) Fetch(ctx context.Context, param company_request.CompanyParam) (res []model.Company, err error) {
	filter(&c.db, param)
	err = c.db.Find(&res).Error
	return
}

func (c companyGorm) Get(ctx context.Context, param company_request.CompanyParam) (res model.Company, err error) {
	filter(&c.db, param)
	err = c.db.First(&res).Error
	return
}

func (c companyGorm) Store(ctx context.Context, company *model.Company) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Create(&company).Error
}

func (c companyGorm) Update(ctx context.Context, company *model.Company, condition ...company_request.CompanyParam) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Updates(&company).Error
}

func (c companyGorm) Delete(ctx context.Context, uuid string) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Delete(model.Company{}, fmt.Sprintf("uuid=%v", uuid)).Error
}
