package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func BranchSeed(db *gorm.DB) (r []model.Branch, err error) {
	err = faker.FakeData(&r)
	if err != nil {
		return
	}
	err = db.Debug().Create(&r).Error
	return
}
