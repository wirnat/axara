package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterServiceDaySeeder(caps []model.Capster, db *gorm.DB) error {
	for _, v := range caps {
		r := random(0, 6)
		for i := 0; i <= r; i++ {
			var sd model.CapsterServiceDay
			err := faker.FakeData(&sd)
			if err != nil {
				return err
			}
			sd.CapsterUUID = v.UUID
			sd.Day = day(i).String()

			err = db.Create(&sd).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
