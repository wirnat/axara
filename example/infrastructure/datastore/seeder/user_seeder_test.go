package seeder

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"gitlab.com/aksaratech/barber-backend/infrastructure/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func connectDB() *gorm.DB {
	env.LoadConf("../../config.yaml")
	m := env.ENV.Pgsql

	dsn, err := pgx.ParseConfig(m.LinkDsn("barber"))
	if err != nil {
		return nil
	}

	sqlDB := stdlib.OpenDB(*dsn)

	//sqlDB, err := sql.Open("postgres", m.LinkDsn("barber"))
	//if err != nil {
	//	fmt.Println("Failed to connect DB")
	//	panic(err)
	//}

	pgsqlConf := postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(pgsqlConf), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return db
}

func TestUserSeed(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{db: connectDB()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := UserSeed(5, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("UserSeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
