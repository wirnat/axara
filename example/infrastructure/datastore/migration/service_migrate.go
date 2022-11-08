package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func ServiceMigrate(db *gorm.DB) error {
	fmt.Println("migrating service module")
	return db.AutoMigrate(
		&model.Category{},
		&model.Service{},
		&model.Slot{},
	)
}
