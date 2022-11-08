package v1

type InfrastructureBuilder struct {
	Constructor
	ModelTrait []ModuleTrait `json:"model_trait" yaml:"model_trait"`
}
