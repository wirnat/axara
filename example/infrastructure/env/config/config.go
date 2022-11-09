package conf

type AppConf struct {
	System              System              `mapstructure:"system" json:"system" yaml:"system"`
	Mysql               Mysql               `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql               Pgsql               `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Redis               Redis               `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha             Captcha             `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	OssLocal            OssLocal            `mapstructure:"oss_local" json:"oss_local" yaml:"oss_local"`
	OssAli              OssAli              `mapstructure:"oss_ali" json:"oss_ali" yaml:"oss_ali"`
	Email               Email               `mapstructure:"email" json:"email" yaml:"email"`
	WaBlast             WaBlast             `mapstructure:"wablast" json:"wablast" yaml:"wablast"`
	WaBotika            WaBotika            `mapstructure:"wabotika" json:"wabotika" yaml:"wabotika"`
	GoogleMap           GoogleMap           `mapstructure:"google_map" json:"google_map" yaml:"google_map"`
	JWT                 JWT                 `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ExternalApiEndpoint ExternalApiEndpoint `mapstructure:"external_api_endpoint" json:"external_api_endpoint" yaml:"external_api_endpoint"`
	KeyCloak            KeyCloak            `mapstructure:"keycloak" json:"keycloak" yaml:"keycloak"`
	Kafka               Kafka               `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Cors                CORS                `mapstructure:"cors" json:"cors" yaml:"cors"`
	OssGoogleStorage    OssGoogleStorage    `mapstructure:"oss_google_storage" json:"oss_google_storage" yaml:"oss_google_storage"`
}
