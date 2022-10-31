package contextor

type Framework interface {
	Run(...string) error
	Get(path string, handler ...func(ctx *Contextor) error)
	Post(path string, handler ...func(ctx *Contextor) error)
	Put(path string, handler ...func(ctx *Contextor) error)
	Delete(path string, handler ...func(ctx *Contextor) error)
	Group(path string, handler ...func(ctx *Contextor) error) Framework
}
