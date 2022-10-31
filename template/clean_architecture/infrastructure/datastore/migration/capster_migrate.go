package migration

import (
	"fmt"
	"gitlab.com/aksaratech/barber-backend/domain/model"
	"gorm.io/gorm"
)

func CapsterMigrate(db *gorm.DB) error {
	fmt.Println("migrating capster module")
	return db.AutoMigrate(
		&model.Capster{},
		&model.CapsterCertificate{},
		&model.CapsterFavorite{},
		&model.CapsterMedia{},
		&model.CapsterProfile{},
		&model.CapsterRating{},
		&model.CapsterRatingMedia{},
		&model.CapsterService{},
		&model.CapsterServiceMedia{},
		&model.CapsterServiceUnavailability{},
		&model.CapsterServiceDay{},
		&model.CapsterSetting{},
		&model.CapsterProduct{},
	)
}
