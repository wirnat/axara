package branch_gorm

import (
	"context"
	"fmt"
	"gitlab.com/wirawirw/aksara-cli/example/model"
	"gitlab.com/wirawirw/aksara-cli/modules/branch/branch_request"
	"gorm.io/gorm"
)

type branchGorm struct {
	db gorm.DB
}

func New(db gorm.DB) *branchGorm {
	return &branchGorm{db: db}
}

func (c branchGorm) Fetch(ctx context.Context, param branch_request.BranchParam) (res []model.Branch, err error) {
	filter(&c.db, param)
	err = c.db.Find(&res).Error
	return
}

func (c branchGorm) Get(ctx context.Context, param branch_request.BranchParam) (res model.Branch, err error) {
	filter(&c.db, param)
	err = c.db.First(&res).Error
	return
}

func (c branchGorm) Store(ctx context.Context, branch *model.Branch) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Create(&branch).Error
}

func (c branchGorm) Update(ctx context.Context, branch *model.Branch, condition ...branch_request.BranchParam) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Updates(&branch).Error
}

func (c branchGorm) Delete(ctx context.Context, uuid string) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return db.Delete(model.Branch{}, fmt.Sprintf("uuid=%v", uuid)).Error
}
