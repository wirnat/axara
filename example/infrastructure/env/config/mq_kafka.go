package conf

type Kafka struct {
	Username         string   `mapstructure:"username" json:"username" yaml:"username"`
	Password         string   `mapstructure:"password" json:"password" yaml:"password"`
	Address          []string `mapstructure:"address" json:"address" yaml:"address"`
	PrefixGroupName  string   `mapstructure:"prefix_group_name" json:"prefix_group_name" yaml:"prefix_group_name"`
	SecurityProtocol string   `mapstructure:"security_protocol" json:"security_protocol" yaml:"security_protocol"`
	SSLCaLocation    string   `mapstructure:"ssl_ca_location" json:"ssl_ca_location" yaml:"ssl_ca_location"`
	SASLMechanism    string   `mapstructure:"sasl_mechanism" json:"sasl_mechanism" yaml:"sasl_mechanism"`
}
