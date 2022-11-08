package env

import (
	conf "gitlab.com/aksaratech/barber-backend/infrastructure/env/config"
)

const (
	ConfigEnv           = "CONFIG_FILE"
	ConfigFile          = "config.yaml"
	MigrationSourcePath = "./infrastructure/migration/source"
	//Message Broker
	MbCustomerTopic = "customer"
	MbMediaTopic    = "media"
	MbCatalogTopic  = "catalog"
)

var ENV conf.AppConf
