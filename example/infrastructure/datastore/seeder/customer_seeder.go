package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CustomerSeeder(db *gorm.DB) (cs []model.Customer, err error) {
	var u []model.UserRelation
	err = db.Preload("Capster").Find(&u).Error
	if err != nil {
		return
	}

	for _, v := range u {
		if v.Capster.ID == 0 {
			var c model.Customer
			err := faker.FakeData(&c)
			if err != nil {
				return nil, err
			}
			c.UserUUID = v.UUID
			err = db.Create(&c).Error
			if err != nil {
				return nil, err
			}
			cs = append(cs, c)
		}
	}

	return
}
