package main

import (
	//@Generate import-main
	"github.com/labstack/echo/v4"
)
func init(){
    //@Generate init
}

func main() {
    //@Generate init-main
	e := echo.New()

	//@Generate route

	e.Start(":8080")
}
