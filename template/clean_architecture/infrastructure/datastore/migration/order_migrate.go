package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func OrderMigrate(db *gorm.DB) error {
	fmt.Println("migrating order module")
	return db.AutoMigrate(
		&model.Order{},
		&model.OrderService{},
		&model.OrderService{},
		&model.OrderSlot{},
	)
}
