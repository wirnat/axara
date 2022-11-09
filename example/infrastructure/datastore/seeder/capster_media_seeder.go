package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterMediaSeed(c []model.Capster, db *gorm.DB) error {
	for _, v := range c {
		r := random(0, 15)
		for i := 0; i < r; i++ {
			var m model.Media
			var cm model.CapsterMedia
			err := faker.FakeData(&m)
			m.UserUUID = v.UserUUID
			if err != nil {
				return err
			}
			err = db.Create(&m).Error
			if err != nil {
				return err
			}
			cm.MediaUUID = m.UUID
			cm.CapsterUUID = v.UUID
			err = db.Create(&cm).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
