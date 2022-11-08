package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func MediaMigrate(db *gorm.DB) error {
	fmt.Println("migrating media module")
	return db.AutoMigrate(&model.Media{})
}
