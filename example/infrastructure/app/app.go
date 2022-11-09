package app

import (
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor"
	"gitlab.com/aksaratech/barber-backend/infrastructure/datastore"
)

type App struct {
	DB        datastore.Database
	Framework contextor.Framework
}
