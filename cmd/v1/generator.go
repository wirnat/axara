package v1

type Generator interface {
	Generate(constructor Constructor) error
}
