package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func ServiceSeed(catSeed, serviceSeed int, db *gorm.DB) (ss []model.Service, err error) {
	cat := make([]string, 0)
	for i := 0; i < catSeed; i++ {
		var c model.Category
		err := faker.FakeData(&c)
		if err != nil {
			return nil, err
		}
		err = db.Create(&c).Error
		if err != nil {
			return nil, err
		}
		cat = append(cat, c.UUID)
	}

	for i := 0; i < serviceSeed; i++ {
		x := random(0, catSeed)
		var s model.Service
		err := faker.FakeData(&s)
		s.CategoryUUID = cat[x]
		if err != nil {
			return nil, err
		}
		err = db.Create(&s).Error
		if err != nil {
			return nil, err
		}

		ss = append(ss, s)
	}
	return
}
