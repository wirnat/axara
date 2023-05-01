package datastore

import (
	"fmt"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadDBGorm() *gorm.DB {
	var db *gorm.DB
	var err error
	dbType := env.ENV.System.DbType
	switch dbType {
	case "mysql":
		m := env.ENV.Mysql
		mysqlConf := mysql.Config{
			DSN:                       m.Dsn(),
			DefaultStringSize:         191,
			SkipInitializeWithVersion: false,
		}
		db, err = gorm.Open(mysql.New(mysqlConf), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			//Logger:                                   logger.Default.LogMode(logger.Info),
			Logger: logger.Default.LogMode(logger.Silent),
		})

		break
	case "pgsql":
		fmt.Println("--initiate post")
		m := env.ENV.Pgsql

		dsn, err := pgx.ParseConfig(m.LinkDsn(env.ENV.Pgsql.DbName))
		if err != nil {
			return nil
		}

		sqlDB := stdlib.OpenDB(*dsn)
		err = sqlDB.Ping()
		if err != nil {
			panic("error connecting db, " + err.Error())
		}

		//sqlDB, err := sql.Open("postgres", m.LinkDsn("barber"))
		//if err != nil {
		//	fmt.Println("Failed to connect DB")
		//	panic(err)
		//}

		pgsqlConf := postgres.Config{
			Conn:                 sqlDB,
			PreferSimpleProtocol: true,
		}

		db, err = gorm.Open(postgres.New(pgsqlConf), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			//Logger:                                   logger.Default.LogMode(logger.Info),
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic("error setup gorm database, " + err.Error())
		}

		con, _ := db.DB()
		err = con.Ping()
		if err != nil {
			panic("database not responding")
		}
		break
	default:
		panic("conf fail")
	}

	if err != nil {
		panic(err)
	}

	//err = migration.MigrateAll(db)
	//err = seeder.SeedAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("==Gorm loaded")
	return db
}
