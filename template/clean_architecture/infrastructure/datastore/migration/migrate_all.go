package migration

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	err := UserMigrate(db)
	err = CapsterMigrate(db)
	err = CustomerMigrate(db)
	err = BranchMigrate(db)
	err = MediaMigrate(db)
	err = OrderMigrate(db)
	err = ServiceMigrate(db)
	err = RegionMigrate(db)
	if err != nil {
		return err
	}
	return nil
}
