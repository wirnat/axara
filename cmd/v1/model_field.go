package v1

type ModelField struct {
	Json  string                 `json:"json"`
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	IsPtr bool                   `json:"is_ptr"`
	Meta  map[string]interface{} `json:"meta"`
}
