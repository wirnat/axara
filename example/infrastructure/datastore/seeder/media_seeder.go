package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func MediaSeed(seed int, db *gorm.DB) error {
	for i := 0; i < seed; i++ {
		var m model.Media
		err := faker.FakeData(&m)
		if err != nil {
			return err
		}
	}

	return nil
}
