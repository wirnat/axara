package branch_gorm

import (
	"gorm.io/gorm"
	"github.com/wirnat/axara/module/branch/request/branch_request"
)

func filter(db *gorm.DB, Param branch_request.BranchParam) {
	if Param.ID != nil {
		*db = *db.Where("id", *Param.ID)
	}
	if Param.UUID != nil {
		*db = *db.Where("uuid", *Param.UUID)
	}
	if Param.CreatedAt != nil {
		*db = *db.Where("created_at", *Param.CreatedAt)
	}
	if Param.UpdatedAt != nil {
		*db = *db.Where("updated_at", *Param.UpdatedAt)
	}
	if Param.DeletedAt != nil {
		*db = *db.Where("deleted_at", *Param.DeletedAt)
	}
	if Param.CompanyID != nil {
		*db = *db.Where("company_id", *Param.CompanyID)
	}
	if Param.Name != nil {
		*db = *db.Where("name", *Param.Name)
	}
	if Param.Description != nil {
		*db = *db.Where("description", *Param.Description)
	}
}
