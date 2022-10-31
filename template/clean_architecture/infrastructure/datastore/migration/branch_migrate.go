package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func BranchMigrate(db *gorm.DB) error {
	fmt.Println("migrating branch module")
	return db.AutoMigrate(&model.Branch{})
}
