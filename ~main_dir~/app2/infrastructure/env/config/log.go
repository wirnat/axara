package config

type Log struct {
	MaxSize   int    `json:"max_size" yaml:"max_size" mapstructure:"max_size"`
	Plugin    string `json:"plugin" yaml:"plugin" mapstructure:"plugin"`
	Compress  bool   `json:"compress" yaml:"compress" mapstructure:"compress"`
	Location  string `json:"location" yaml:"location" mapstructure:"location"`
	MaxAge    int    `json:"max_age" yaml:"max_age" mapstructure:"max_age"`
	MaxBackup int    `json:"max_backup" yaml:"max_backup" mapstructure:"max_backup"`
}
