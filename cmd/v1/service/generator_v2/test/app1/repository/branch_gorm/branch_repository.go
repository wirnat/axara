package branch_gorm

import (
	"context"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/model"
    "<no value>"
    "<no value>"
	"<no value>"
	"gorm.io/gorm"
)

type branchGorm struct {
	db gorm.DB
}

func New(db gorm.DB) *branchGorm {
    db.AutoMigrate(&model.Branch{})
	return &branchGorm{db: db}
}

func (c branchGorm) Fetch(ctx context.Context, Param branch_request.BranchParam) (res []model.Branch, err error) {
	filter(&c.db, Param)
	err = c.db.Find(&res).Error
	return
}

func (c branchGorm) Get(ctx context.Context, Param branch_request.BranchParam) (res model.Branch, err error) {
	filter(&c.db, Param)
	err = c.db.First(&res).Error
	return
}

func (c branchGorm) Store(ctx context.Context, branch *model.Branch) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	return c.db.Create(&branch).Error
}

func (c branchGorm) Update(ctx context.Context, branch *model.Branch, condition ...branch_request.BranchParam) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}

	for _, p := range condition {
        filter(&c.db, p)
    }

	return c.db.Updates(&branch).Error
}

func (c branchGorm) Delete(ctx context.Context, uuid string) error {
	db, err := getTx(ctx)
	if err == nil {
		c.db = *db
	}
	return c.db.Debug().Where("uuid=?", uuid).Delete(&model.Branch{}).Error
}

func (c branchGorm) Paginate(ctx context.Context, param branch_request.BranchParam) (response.BranchPaginate, error) {
	res := response.BranchPaginate{
		Pagination: &paginator.Pagination{
			PaginateReq: param.PaginateReq,
		},
	}

	filter(&c.db, param)
	err := c.db.Debug().Scopes(paginator.PaginateV2(&res.Result, &c.db, res.Pagination)).Find(&res.Result).Error
	if err != nil {
		return response.BranchPaginate{}, err
	}

	return res, nil
}

