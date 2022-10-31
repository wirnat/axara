package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterCertifiaceSeed(c []model.Capster, db *gorm.DB) error {
	for _, v := range c {
		r := random(0, 10)
		if r > 5 {
			x := random(1, 5)
			for i := 0; i < x; i++ {
				var cert model.CapsterCertificate
				var media model.Media
				err := faker.FakeData(&cert)
				media.UserUUID = v.UserUUID
				err = faker.FakeData(&media)
				err = db.Create(&media).Error
				if err != nil {
					return err
				}
				if err != nil {
					return err
				}
				cert.CapsterUUID = v.UUID
				cert.MediaUUID = media.UUID
				err = db.Create(&cert).Error
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
