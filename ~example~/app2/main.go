package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	//@Generate route

	e.Start(":8080")
}
