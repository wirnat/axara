package config

import "fmt"

type Pgsql struct {
	Path              string `mapstructure:"path" json:"path" yaml:"path" `
	Port              string `mapstructure:"port" json:"port" yaml:"port"`
	Config            string `mapstructure:"config" json:"config" yaml:"config"`
	DbName            string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	Username          string `mapstructure:"username" json:"username" yaml:"username"`
	Password          string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection" json:"max_idle_connection" yaml:"max_idle_connection"`
	MaxOpenConnection int    `mapstructure:"max_open_connection" json:"max_open_connection" yaml:"max_open_connection"`
	LogMode           string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	LogZap            bool   `mapstructure:"log_zap" json:"log_zap" yaml:"log_zap"`
}

func (p *Pgsql) Dsn() string {
	if p.Port == "" {
		p.Port = "5432"
	}
	return fmt.Sprintf("host=%s user=%s password=%s port=%s %s", p.Path, p.Username, p.Password, p.Port, p.Config)
}

func (p Pgsql) LinkDsn(dbName string) string {
	if p.Port == "" {
		p.Port = "5432"
	}
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s %s", p.Path, p.Username, dbName, p.Port,
		p.Password, p.Config)
}
