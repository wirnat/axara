package seeder

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func UserSeed(user int, db *gorm.DB) (users []model.User, err error) {
	for i := 0; i < user; i++ {
		var u model.User
		err := faker.FakeData(&u)
		if err != nil {
			return nil, err
		}

		u.ID = 0

		err = db.Create(&u).Error
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error insert user")
		}
		users = append(users, u)
	}

	return
}
