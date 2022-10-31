package conf

import "strconv"

type MqRabbitMq struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (r MqRabbitMq) ConnectionString() string {
	return r.Address + ":" + strconv.Itoa(r.Port)
}
