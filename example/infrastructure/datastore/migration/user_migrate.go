package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func UserMigrate(db *gorm.DB) error {
	fmt.Println("migrating user module")
	return db.AutoMigrate(&model.User{}, &model.UserBranch{}, &model.UserCompany{})
}
