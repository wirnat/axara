package company_gorm

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"gorm.io/gorm"
)

type companyGorm struct {
	db gorm.DB
}

func NewCompanyGorm(db gorm.DB) *companyGorm {
	return &companyGorm{db: db}
}

func (b companyGorm) Store(ctx context.Context, company *model.Company) error {
	return b.db.Create(company).Error
}

func (b companyGorm) Update(ctx context.Context, company *model.Company) error {
	return b.db.Updates(company).Error
}

func (b companyGorm) Get(ctx context.Context, uuid string) (r model.Company, err error) {
	err = b.db.Where("uuid=?", uuid).First(&r).Error
	return
}

func (b companyGorm) Fetch(ctx context.Context) (r []model.Company, err error) {
	err = b.db.First(&r).Error
	return
}
