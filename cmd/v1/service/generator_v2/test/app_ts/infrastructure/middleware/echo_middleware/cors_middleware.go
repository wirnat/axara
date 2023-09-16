package echo_middleware

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app_ts/infrastructure/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AllowCors() echo.MiddlewareFunc {
	var corsEnv = env.ENV.Cors

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     corsEnv.AllowOrigins,
		AllowOriginFunc:  nil,
		AllowMethods:     corsEnv.AllowMethods,
		AllowHeaders:     corsEnv.AllowHeaders,
		AllowCredentials: corsEnv.AllowCredentials,
		ExposeHeaders:    corsEnv.ExposeHeaders,
		MaxAge:           0,
	})
}
