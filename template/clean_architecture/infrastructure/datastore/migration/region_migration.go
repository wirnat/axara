package migration

import (
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func RegionMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Village{})
	err = db.AutoMigrate(&model.District{})
	err = db.AutoMigrate(&model.Province{})
	err = db.AutoMigrate(&model.Nationality{})
	return err
}
