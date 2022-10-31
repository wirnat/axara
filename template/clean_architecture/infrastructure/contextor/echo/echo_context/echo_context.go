package echo_context

import (
	"context"
	"github.com/labstack/echo/v4"
	"gitlab.com/aksaratech/barber-backend/infrastructure/contextor"
	"mime/multipart"
	"net/http"
)

type echoContext struct {
	Echo echo.Context
}

func (e echoContext) Validate(i interface{}) error {
	return e.Echo.Validate(i)
}

func (e echoContext) FormFile(name string) (*multipart.FileHeader, error) {
	return e.Echo.FormFile(name)
}

func (e echoContext) ToContext() (c context.Context) {
	user := e.Echo.Get(contextor.UserID)
	c = context.WithValue(e.Echo.Request().Context(), contextor.UserID, user)
	c = context.WithValue(c, contextor.Company, e.Echo.Get(contextor.Company))
	c = context.WithValue(c, contextor.Token, e.Echo.Get(contextor.Token))
	return
}

func (e echoContext) Request() *http.Request {
	return e.Echo.Request()
}

func (e echoContext) QueryParam(name string) string {
	return e.Echo.QueryParam(name)
}

func (e echoContext) JSON(code int, i interface{}) error {
	return e.Echo.JSON(code, i)
}

func (e echoContext) Bind(dest interface{}) error {
	return e.Echo.Bind(dest)
}

func (e echoContext) Param(name string) string {
	return e.Echo.Param(name)
}

func NewEchoContext(echo echo.Context) *echoContext {
	return &echoContext{Echo: echo}
}
