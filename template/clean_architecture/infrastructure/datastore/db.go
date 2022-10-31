package datastore

import (
	"database/sql"
	"github.com/aksara-tech/aksarabase"
	"gorm.io/gorm"
)

type Database struct {
	GORM  *gorm.DB
	ADB   *aksarabase.ADB
	SQLDB *sql.DB
}
