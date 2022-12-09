package v1

type ModuleTrait struct {
	Name     string   `json:"name" yaml:"name"`
	Dir      string   `json:"dir" yaml:"dir"`
	FileName string   `json:"file_name" yaml:"file_name"`
	Template string   `json:"template" yaml:"template"`
	Active   bool     `json:"active" yaml:"active"`
	CMD      []string `json:"cmd" yaml:"CMD"`
}
