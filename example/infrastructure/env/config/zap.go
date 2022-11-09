package conf

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	ShowLine      bool   `mapstructure:"show_line" json:"show_line" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`
	StackTraceKey string `mapstructure:"stack_trace_key" json:"stack_trace_key" yaml:"stack_trace_key"`
	LogInConsole  bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"`
}
