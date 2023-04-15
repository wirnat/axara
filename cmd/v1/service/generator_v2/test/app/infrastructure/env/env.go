package env

import conf "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/env/conf"

const (
	ConfigEnv           = "CONFIG_FILE"
	ConfigFile          = "config.yaml"
	MigrationSourcePath = "./infrastructure/migration/source"
	//Message Broker
	MbCustomerTopic = "customer"
	MbMediaTopic    = "media"
	MbCatalogTopic  = "catalog"
	DefaultLogPath  = "./log"
	AppModeDev  = "dev"
    AppModeProd = "production"
)

var ENV conf.AppConf
