package conf

type ExternalApiEndpoint struct {
	LotteryUrl          string `mapstructure:"lottery_url" json:"lottery_url" yaml:"lottery_url"`
	MsSecureUrl         string `mapstructure:"ms_secure_url" json:"ms_secure_url" yaml:"ms_secure_url"`
	MsSecureUrlFrontend string `mapstructure:"ms_secure_url_frontend" json:"ms_secure_url_frontend" yaml:"ms_secure_url_frontend"`
	MsProductUrl        string `mapstructure:"ms_product_url" json:"ms_product_url" yaml:"ms_product_url"`
	MsAggregationUrl    string `mapstructure:"ms_aggregation_url" json:"ms_aggregation_url" yaml:"ms_aggregation_url" `
}
