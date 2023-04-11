package testing_env

import "github.com/labstack/echo/v4"

func Route() {
	//@Generate app
	e := echo.New()

	//@Generate route

	//@Generate endApp
	e.Start(":9999")
}
