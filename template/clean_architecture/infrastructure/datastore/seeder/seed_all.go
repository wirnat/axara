package seeder

import (
	"fmt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	customGenerator()
	fmt.Println("seed user")
	us, err := UserSeed(200, db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster")
	cs, err := CapsterSeed(us, db)
	if err != nil {
		return err
	}
	fmt.Println("seed service")
	se, err := ServiceSeed(5, 10, db)
	if err != nil {
		return err
	}
	fmt.Println("seed customer")
	cu, err := CustomerSeeder(db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster certificate")
	err = CapsterCertifiaceSeed(cs, db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster media")
	err = CapsterMediaSeed(cs, db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster favorite")
	err = CapsterFavoriteSeeder(cs, cu, db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster rating")
	err = CapsterRatingSeeder(cs, cu, db)
	if err != nil {
		return err
	}
	fmt.Println("seed capster service")
	err = CapsterServiceSeeder(cs, se, db)
	if err != nil {
		return err
	}
	err = CapsterServiceDaySeeder(cs, db)
	fmt.Println("seed capster service")
	_, err = BranchSeed(db)
	if err != nil {
		return err
	}
	//SeedingRegionStructure(db)
	return nil
}
