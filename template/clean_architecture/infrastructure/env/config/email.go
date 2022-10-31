package conf

type Email struct {
	Host       string `mapstructure:"host" json:"host" yaml:"host"`
	Port       string `mapstructure:"port" json:"port" yaml:"port"`
	Username   string `mapstructure:"username" json:"username" yaml:"username"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
	Encryption string `mapstructure:"encryption" json:"encryption" yaml:"encryption"`
	From       string `mapstructure:"from" json:"from" yaml:"from"`
}
