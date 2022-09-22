package cmd

type Trait struct {
	Meta map[string]interface{} `json:"meta"`
	ModelTrait
	ModuleTrait
}
