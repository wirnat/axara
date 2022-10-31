package conf

type GoogleMap struct {
	ApiKeyToken string `mapstructure:"key_token" json:"key_token" yaml:"key_token"`
	ClientID    string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	Signature   string `mapstructure:"signature" json:"signature" yaml:"signature"`
}
