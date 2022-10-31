package conf

type AppCatalogService struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	HttpPort string `json:"http_port" yaml:"http_port"`
	GrpcPort string `json:"grpc_port" yaml:"grpc_port"`
	IsSSL    bool   `mapstructure:"isSSL" json:"isSSL" yaml:"isSSL"`
}
type AppMediaService struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	HttpPort string `json:"http_port" yaml:"http_port"`
	GrpcPort string `json:"grpc_port" yaml:"grpc_port"`
	IsSSL    bool   `mapstructure:"isSSL" json:"isSSL" yaml:"isSSL"`
}
