package conf

type System struct {
	Env                 string `mapstructure:"env" json:"env" yaml:"env"`
	Address             string `mapstructure:"address" json:"address" yaml:"address"`
	OssType             string `mapstructure:"oss_type" json:"oss_type" yaml:"oss_type"`
	DbType              string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
	MessageBrokerType   string `mapstructure:"message_broker_type" json:"message_broker_type" yaml:"message_broker_type"`
	MigrationSourcePath string `mapstructure:"migration_source_path" json:"migration_source_path" yaml:"migration_source_path"`
	CsPhone             string `mapstructure:"cs_phone" json:"cs_phone" yaml:"cs_phone"`
	CsEmail             string `mapstructure:"cs_email" json:"cs_email" yaml:"cs_email"`
	LogLocation         string `mapstructure:"log_location" json:"log_location" yaml:"log_location"`
}
