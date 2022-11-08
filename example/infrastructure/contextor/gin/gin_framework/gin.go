package gin_framework

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor/gin/gin_context"
)

type ginx struct {
	Gin *gin.Engine
}

func (g *ginx) Put(path string, handler ...func(ctx *contextor.Contextor) error) {
	g.Gin.PUT(path, func(ctx *gin.Context) {
		for _, h := range handler {
			context := g.convert(ctx)
			err := h(context)
			if err != nil {
				ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				return
			}
		}
	})
}

func (g *ginx) Delete(path string, handler ...func(ctx *contextor.Contextor) error) {
	g.Gin.DELETE(path, func(ctx *gin.Context) {
		for _, h := range handler {
			context := g.convert(ctx)
			err := h(context)
			if err != nil {
				if err != nil {
					ctx.JSON(500, contextor.Response{
						Code: 500,
						Data: err,
						Msg:  err.Error(),
					})
					return
				}
			}
		}
	})
}

func (g *ginx) Group(path string, handler ...func(ctx *contextor.Contextor) error) contextor.Framework {
	return &contextor.Group{
		Framework:  g,
		Prefix:     path,
		PrefixFunc: handler,
	}
}

func (g *ginx) Run(s ...string) error {
	return g.Gin.Run(s...)
}

func (g *ginx) Post(path string, handler ...func(ctx *contextor.Contextor) error) {
	g.Gin.POST(path, func(ctx *gin.Context) {
		for _, h := range handler {
			context := g.convert(ctx)
			err := h(context)
			if err != nil {
				ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				return
			}
		}
	})
}

func (g *ginx) Get(path string, handler ...func(ctx *contextor.Contextor) error) {
	g.Gin.GET(path, func(ctx *gin.Context) {
		for _, h := range handler {
			context := g.convert(ctx)
			err := h(context)
			if err != nil {
				ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				return
			}
		}
	})
}

func NewGinx(gin *gin.Engine) *ginx {
	return &ginx{Gin: gin}
}

func (g *ginx) convert(ctx *gin.Context) *contextor.Contextor {
	c := gin_context.NewGinContext(ctx)
	return contextor.NewContextor(c)
}
