package v1

type Constructor struct {
	GitAccessKey  string                       `json:"git_access_key" yaml:"git_access_key"`
	Key           string                       `json:"key" yaml:"key"`
	Lang          string                       `json:"lang" yaml:"lang"`
	ModelPath     string                       `json:"model_path" yaml:"model_path"`
	ModuleName    string                       `json:"module_name" yaml:"module_name"`
	Jobs          []Job                        `json:"jobs" yaml:"jobs"`
	Meta          map[string]string            `json:"meta" yaml:"meta"`
	IncludeJobs   []string                     `json:"include_jobs"  yaml:"include_jobs"`
	IncludeTraits []string                     `json:"include_traits" yaml:"include_traits"`
	Models        map[string]map[string]string `json:"models" yaml:"models"`
}
