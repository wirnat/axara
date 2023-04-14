package config

type AppConf struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
	//TODO: You can add other config struct here and it will auto generate new config yaml
}
