package conf

type Captcha struct {
	KeyLong     int `mapstructure:"key_long" json:"key_long" yaml:"key_long" `
	ImageWidth  int `mapstructure:"image_width" json:"image_width" yaml:"image_width"`
	ImageHeight int `mapstructure:"image_height" json:"image_height" yaml:"image_height"`
}
