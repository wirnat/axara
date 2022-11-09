package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterRatingSeeder(c []model.Capster, u []model.Customer, db *gorm.DB) error {
	x := random(0, len(u)-1)
	for _, v := range c {
		var ins []int
		r := random(0, 10)
		if r > 3 {
			for {
				user := u[x]
				x = random(x, len(u)-1)
				if inserted(x, ins) {
					continue
				}
				if x == len(u)-1 || x > len(u)-3 {
					break
				}
				var cf model.CapsterRating
				err := faker.FakeData(&cf)
				if err != nil {
					return err
				}
				cf.CapsterUUID = v.UUID
				cf.CustomerUUID = user.UUID
				db.Create(&cf)
				x = random(x, len(u)-1)
				if x == len(u)-1 || x > len(u)-3 {
					break
				}
				ins = append(ins, x)
			}
		}
	}
	return nil
}
