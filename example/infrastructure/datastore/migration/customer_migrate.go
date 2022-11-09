package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CustomerMigrate(db *gorm.DB) error {
	fmt.Println("migrating customer module")
	return db.AutoMigrate(&model.Customer{})
}
