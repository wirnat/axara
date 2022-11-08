package conf

type WaBotika struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	ApiToken string `mapstructure:"token" json:"token" yaml:"token"`
}
