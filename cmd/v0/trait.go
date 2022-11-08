package v0

type Trait struct {
	Meta map[string]interface{} `json:"meta"`
	ModelTrait
	ModuleTrait
}
