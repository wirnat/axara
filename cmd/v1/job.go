package v1

type Job struct {
	Name          string   `json:"name" yaml:"name"`
	Dir           string   `json:"dir" yaml:"dir"`
	FileName      string   `json:"file_name" yaml:"file_name"`
	Template      string   `json:"template" yaml:"template"`
	Active        bool     `json:"active" yaml:"active"`
	CMD           []string `json:"cmd" yaml:"CMD"`
	Tags          []string `json:"tags" yaml:"tags"`
	GenerateIn    string   `json:"generate_in" yaml:"generate_in"`
	SingleExecute bool     `json:"single_execute" yaml:"single_execute"`
}
