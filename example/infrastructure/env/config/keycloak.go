package conf

type KeyCloak struct {
	KeyCloakIP   string `mapstructure:"key_cloak_ip" json:"key_cloak_ip" yaml:"key_cloak_ip"`
	Realm        string `mapstructure:"realm" json:"realm" yaml:"realm"`
	ClientID     string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	ClientUUID   string `mapstructure:"client_uuid" json:"client_uuid" yaml:"client_uuid"`
	ClientSecret string `mapstructure:"client_secret" json:"client_secret" yaml:"client_secret"`
}
