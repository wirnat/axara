package branch_gorm

import (
	"github.com/wirnat/aksara-cli/modules/branch/branch_request"
	"gorm.io/gorm"
)

func filter(db *gorm.DB, param branch_request.BranchParam) {
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
	if param.CompanyID != nil {
		db = db.Where("CompanyID", *param.CompanyID)
	}
	if param.Name != nil {
		db = db.Where("Name", *param.Name)
	}
	if param.Description != nil {
		db = db.Where("Description", *param.Description)
	}
	if param.Email != nil {
		db = db.Where("Email", *param.Email)
	}
	if param.Phone != nil {
		db = db.Where("Phone", *param.Phone)
	}
	if param.PicName != nil {
		db = db.Where("PicName", *param.PicName)
	}
	if param.PicPhone != nil {
		db = db.Where("PicPhone", *param.PicPhone)
	}
	if param.PicEmail != nil {
		db = db.Where("PicEmail", *param.PicEmail)
	}
	if param.Address != nil {
		db = db.Where("Address", *param.Address)
	}
	if param.Status != nil {
		db = db.Where("Status", *param.Status)
	}
	if param.VerifiedStatus != nil {
		db = db.Where("VerifiedStatus", *param.VerifiedStatus)
	}
	if param.OpenStatus != nil {
		db = db.Where("OpenStatus", *param.OpenStatus)
	}
	if param.ProfileImageID != nil {
		db = db.Where("ProfileImageID", *param.ProfileImageID)
	}
	if param.OpenedAt != nil {
		db = db.Where("OpenedAt", *param.OpenedAt)
	}
	if param.ClosedAt != nil {
		db = db.Where("ClosedAt", *param.ClosedAt)
	}
	if param.Latitude != nil {
		db = db.Where("Latitude", *param.Latitude)
	}
	if param.Longitude != nil {
		db = db.Where("Longitude", *param.Longitude)
	}
}
