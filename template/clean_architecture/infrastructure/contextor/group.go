package contextor

import (
	"fmt"
)

type Group struct {
	Framework
	Prefix     string
	PrefixFunc []func(ctx *Contextor) error
}

func (g Group) Post(path string, handler ...func(ctx *Contextor) error) {
	handlers := g.PrefixFunc
	handlers = append(handlers, handler...)
	path = fmt.Sprintf("%v%v", g.Prefix, path)
	g.Framework.Post(path, handlers...)
}

func (g Group) Delete(path string, handler ...func(ctx *Contextor) error) {
	handlers := g.PrefixFunc
	handlers = append(handlers, handler...)
	path = fmt.Sprintf("%v%v", g.Prefix, path)
	g.Framework.Delete(path, handlers...)
}

func (g Group) Put(path string, handler ...func(ctx *Contextor) error) {
	handlers := g.PrefixFunc
	handlers = append(handlers, handler...)
	path = fmt.Sprintf("%v%v", g.Prefix, path)
	g.Framework.Put(path, handlers...)
}

func (g Group) Get(path string, handler ...func(ctx *Contextor) error) {
	handlers := g.PrefixFunc
	handlers = append(handlers, handler...)

	path = fmt.Sprintf("%v%v", g.Prefix, path)
	g.Framework.Get(path, handlers...)
}

func (g Group) Group(path string, handler ...func(ctx *Contextor) error) Framework {
	g.Prefix = fmt.Sprintf("%v%v", g.Prefix, path)
	g.PrefixFunc = append(g.PrefixFunc, handler...)

	return g
}
