package contextor

import (
	"context"
	"mime/multipart"
	"net/http"
)

type Context interface {
	Bind(dest interface{}) error
	Param(name string) string
	JSON(code int, i interface{}) error
	QueryParam(name string) string
	ToContext() context.Context
	FormFile(name string) (*multipart.FileHeader, error)
	Validate(interface{}) error
	Request() *http.Request
}
