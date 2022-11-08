package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterServiceSeeder(c []model.Capster, s []model.Service, db *gorm.DB) error {
	for _, v := range c {
		var ins []int
		r := random(0, len(s)-1)
		for {
			ser := s[r]
			r = random(r, len(s)-1)
			if inserted(r, ins) {
				continue
			}
			if r == len(s)-1 || r > len(s)-3 {
				break
			}
			var cs model.CapsterService
			err := faker.FakeData(&cs)
			if err != nil {
				return err
			}
			cs.CapsterUUID = v.UUID
			cs.ServiceUUID = ser.UUID

			err = db.Create(&cs).Error
			if err != nil {
				return err
			}

			ins = append(ins, r)
		}
	}
	return nil
}
