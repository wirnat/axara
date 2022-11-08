package conf

type OssGoogleStorage struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	ProjectID       string `mapstructure:"project_id" json:"project_id" yaml:"project_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name" json:"bucket_name" yaml:"bucket_name"`
	BucketUrl       string `mapstructure:"bucket_url" json:"bucket_url" yaml:"bucket_url"`
	BasePath        string `mapstructure:"base_path" json:"base_path" yaml:"base_path"`
}
