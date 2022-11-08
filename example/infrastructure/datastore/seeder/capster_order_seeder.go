package seeder

import (
	"github.com/bxcodec/faker/v3"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterOrderSeeder(caps model.Capster, cu []model.Customer, db *gorm.DB) error {
	for _, v := range cu {
		var o model.Order
		err := faker.FakeData(&o)
		if err != nil {
			return err
		}
		o.CapsterUUID = caps.UUID
		o.CustomerUUID = v.UUID
		o.BranchUUID = *caps.BranchUUID

		err = db.Create(&o).Error
		if err != nil {
			return err
		}
	}
	return nil
}
