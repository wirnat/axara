package conf

type AppConf struct {
	System              System              `mapstructure:"system" json:"system" yaml:"system"`
	Mysql               Mysql               `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql               Pgsql               `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Log                 Log                 `mapstructure:"log" json:"log" yaml:"log"`
	Cors                CORS                `mapstructure:"cors" json:"cors" yaml:"cors"`
}
