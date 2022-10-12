package simple

import (
	"context"
	"github.com/wirnat/aksara-cli/example/model"
	"gorm.io/gorm"
)

type branchGorm struct {
	db gorm.DB
}

func NewBranchGorm(db gorm.DB) *branchGorm {
	return &branchGorm{db: db}
}

func (b branchGorm) Store(ctx context.Context, branch *model.Branch) error {
	return b.db.Create(branch).Error
}

func (b branchGorm) Update(ctx context.Context, branch *model.Branch) error {
	return b.db.Updates(branch).Error
}

func (b branchGorm) Get(ctx context.Context, uuid string) (r model.Branch, err error) {
	err = b.db.Where("uuid=?", uuid).First(&r).Error
	return
}

func (b branchGorm) Fetch(ctx context.Context) (r []model.Branch, err error) {
	err = b.db.First(&r).Error
	return
}
