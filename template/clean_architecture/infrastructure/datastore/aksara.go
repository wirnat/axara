package datastore

import (
	"github.com/aksara-tech/aksarabase"
	"github.com/aksara-tech/aksarabase/domain"
	"github.com/aksara-tech/aksarabase/domain/constanta"
	"github.com/aksara-tech/aksarabase/usecase/query_executor/debug_executor"
	"gitlab.com/aksaratech/barber-backend/infrastructure/env"
)

func LoadAksaraBase() *aksarabase.ADB {
	m := env.ENV.Mysql
	db := aksarabase.Open(constanta.MYSQL, m.Dsn(), domain.Config{
		Engine: domain.Engine{
			OutputScanner:      nil,
			InputScanner:       nil,
			PointerScanner:     nil,
			InsertQueryBuilder: nil,
			UpdateQueryBuilder: nil,
			SelectBuilder:      nil,
			SqlExecutor:        debug_executor.NewExecutor(constanta.MYSQL, m.Dsn()),
		},
	})

	return db
}
