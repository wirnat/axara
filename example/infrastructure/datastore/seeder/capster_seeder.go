package seeder

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterSeed(users []model.User, db *gorm.DB) (cs []model.Capster, err error) {
	for _, v := range users {
		r := random(0, 5)
		if r > 3 {
			var c model.Capster
			err = faker.FakeData(&c)
			if err != nil {
				return
			}
			c.UserUUID = v.UUID
			err = db.Create(&c).Error
			if err != nil {
				fmt.Println(err)
				return nil, errors.New("error insert capster")
			}

			x := random(0, 5)
			if x > 2 {
				var p model.CapsterProfile
				err = faker.FakeData(&p)
				if err != nil {
					return
				}
				p.CapsterUUID = c.UUID
				err = db.Create(&p).Error
				if err != nil {
					fmt.Println(err)
					return nil, errors.New("error insert capster profile")
				}
			}
			cs = append(cs, c)
		}
	}

	return
}
