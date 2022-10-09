package company_gorm

import (
	"github.com/wirnat/aksara-cli/modules/company/company_request"
	"gorm.io/gorm"
)

func filter(db *gorm.DB, param company_request.CompanyParam) {
	if param.ID != nil {
    db = db.Where("ID", *param.ID)
    }
	if param.UUID != nil {
    db = db.Where("UUID", *param.UUID)
    }
	if param.CreatedAt != nil {
    db = db.Where("CreatedAt", *param.CreatedAt)
    }
	if param.UpdatedAt != nil {
    db = db.Where("UpdatedAt", *param.UpdatedAt)
    }
	if param.DeletedAt != nil {
    db = db.Where("DeletedAt", *param.DeletedAt)
    }
	if param.Name != nil {
    db = db.Where("Name", *param.Name)
    }
	if param.Phone != nil {
    db = db.Where("Phone", *param.Phone)
    }
	if param.Email != nil {
    db = db.Where("Email", *param.Email)
    }
	if param.Description != nil {
    db = db.Where("Description", *param.Description)
    }
	if param.LogoID != nil {
    db = db.Where("LogoID", *param.LogoID)
    }
	if param.Latitude != nil {
    db = db.Where("Latitude", *param.Latitude)
    }
	if param.Longitude != nil {
    db = db.Where("Longitude", *param.Longitude)
    }
}
