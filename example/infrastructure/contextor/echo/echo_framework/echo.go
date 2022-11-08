package echo_framework

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor/echo/echo_context"
	"gitlab.com/aksaratech/barber-backend/infrastructure/validator"
)

type echox struct {
	Echo *echo.Echo
}

func (c *echox) Put(path string, handler ...func(ctx *contextor.Contextor) error) {
	c.Echo.PUT(path, func(ctx echo.Context) error {
		for _, h := range handler {
			context := c.converter(ctx)
			err := h(context)
			if err != nil {
				err := ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (c *echox) Delete(path string, handler ...func(ctx *contextor.Contextor) error) {
	c.Echo.DELETE(path, func(ctx echo.Context) error {
		for _, h := range handler {
			context := c.converter(ctx)
			err := h(context)
			if err != nil {
				err := ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (c *echox) Group(path string, handler ...func(ctx *contextor.Contextor) error) contextor.Framework {
	return &contextor.Group{
		Framework:  c,
		Prefix:     path,
		PrefixFunc: handler,
	}
}

func (c *echox) Get(path string, handler ...func(ctx *contextor.Contextor) error) {
	c.Echo.GET(path, func(ctx echo.Context) error {
		for _, h := range handler {
			context := c.converter(ctx)
			err := h(context)
			if err != nil {
				err := ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

}

func (c *echox) Post(path string, handler ...func(ctx *contextor.Contextor) error) {
	c.Echo.POST(path, func(ctx echo.Context) error {
		for _, h := range handler {
			context := c.converter(ctx)
			err := h(context)
			if err != nil {
				err := ctx.JSON(500, contextor.Response{
					Code: 500,
					Data: err,
					Msg:  err.Error(),
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (c *echox) Run(address ...string) error {
	c.Echo.HideBanner = true
	err := c.Echo.Start(address[0])
	if err != nil {
		panic(err)
	}

	return err
}

func (c *echox) converter(ctx echo.Context) *contextor.Contextor {
	e := echo_context.NewEchoContext(ctx)
	return contextor.NewContextor(e)
}

func NewEchox(echo *echo.Echo) *echox {
	validator.SetupValidation(echo)
	return &echox{Echo: echo}
}
