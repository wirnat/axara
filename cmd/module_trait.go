package cmd

type ModuleTrait struct {
	Name     string                 `json:"name"`
	Dir      string                 `json:"dir"`
	FileName string                 `json:"file_name"`
	Template string                 `json:"template"`
	Meta     map[string]interface{} `json:"meta"`
}
