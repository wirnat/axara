package v1

type InfrastructureBuilder struct {
	Constructor
	ModelTrait []Job `json:"model_trait" yaml:"model_trait"`
}
