package company_gorm

import (
	"gorm.io/gorm"
    "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/domain/request/company_request"
)

func filter(db *gorm.DB, Param company_request.CompanyParam) {
	if Param.ID != nil {
    *db = *db.Where("id", *Param.ID)
    }
	if Param.UUID != nil {
    *db = *db.Where("uuid", *Param.UUID)
    }
	if Param.createdAt != nil {
    *db = *db.Where("created_at", *Param.createdAt)
    }
	if Param.createdBy != nil {
    *db = *db.Where("created_by", *Param.createdBy)
    }
	if Param.updatedAt != nil {
    *db = *db.Where("updated_at", *Param.updatedAt)
    }
	if Param.updatedBy != nil {
    *db = *db.Where("updated_by", *Param.updatedBy)
    }
	if Param.deletedAt != nil {
    *db = *db.Where("deleted_at", *Param.deletedAt)
    }
	if Param.deletedBy != nil {
    *db = *db.Where("deleted_by", *Param.deletedBy)
    }
	if Param.name != nil {
    *db = *db.Where("name", *Param.name)
    }
	if Param.parent != nil {
    *db = *db.Where("parent", *Param.parent)
    }
	if Param.parentUUID != nil {
    *db = *db.Where("parent_uuid", *Param.parentUUID)
    }
	if Param.subCompany != nil {
    *db = *db.Where("sub_company", *Param.subCompany)
    }
	if Param.type != nil {
    *db = *db.Where("type", *Param.type)
    }
}
