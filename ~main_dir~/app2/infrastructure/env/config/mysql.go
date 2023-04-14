package config

import "fmt"

type Mysql struct {
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

func (m *Mysql) Dsn() string {
	if m.Port == "" {
		m.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Path, m.Port, m.DbName, m.Config)
}
