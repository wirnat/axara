package conf

type JWT struct {
	SigningKey  string `mapstructure:"signing_key" json:"signingKey" yaml:"signing_key"`
	ExpiresTime int64  `mapstructure:"expires_time" json:"expiresTime" yaml:"expires_time"`
	BufferTime  int64  `mapstructure:"buffer_time" json:"bufferTime" yaml:"buffer_time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
