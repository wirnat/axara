package paginator

import (
	"context"
	"gorm.io/gorm"
	"math"
	"strconv"
)

func PaginateV2(value interface{}, dbT *gorm.DB, pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	var totalRows int64
	dbT.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func PaginateV1(ctx context.Context) func(db *gorm.DB) *gorm.DB {
	page := ctx.Value("page").(string)
	pageSize := ctx.Value("page_size").(string)
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSize)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}
