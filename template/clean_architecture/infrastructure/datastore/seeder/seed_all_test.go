package seeder

import (
	"gitlab.com/aksaratech/barber-backend/infrastructure/datastore"
	"gitlab.com/aksaratech/barber-backend/infrastructure/env"
	"gorm.io/gorm"
	"testing"
)

func TestSeedAll(t *testing.T) {
	env.LoadConf("../../../config.yaml")

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test seeder",
			args:    args{db: datastore.LoadDBGorm()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SeedAll(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("SeedAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
