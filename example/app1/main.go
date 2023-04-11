package main

import "github.com/labstack/echo/v4"

func main() {
	//@Generate app
	e := echo.New()

	//@Generate route

	//@Generate endApp
	e.Start(":8080")

}
