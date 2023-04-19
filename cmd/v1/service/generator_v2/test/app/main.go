package main

import (
	"github.com/labstack/echo/v4"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/env"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/logger"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/cmd"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/datastore"
	_ "github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/docs"
	"fmt"
)

//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init -g main.go -o docs

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
// 	@securityDefinitions.apikey ApiKeyAuth
// 	@in header
// 	@name Authorization
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func init() {
	env.LoadConf()
	logger.SetupLogger("storage/log")
}

func main() {
	db := datastore.LoadDBGorm()
	e := echo.New()
	e.Use(echo_middleware.AllowCors())
	e.HideBanner=true
	app := cmd.NewMyApp(*e, *db)
	app.Init()
    for _, route := range e.Routes() {
    	fmt.Printf("%s\t::\t%s\n", route.Method, route.Path)
    }
    initSwagger(e)
  	e.Start(env.ENV.System.Address)
}

func initSwagger(engine *echo.Echo) {
	docs.SwaggerInfo.Title = "My App"
	docs.SwaggerInfo.Description = "Description App"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = env.ENV.System.Address
	//docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	engine.GET("doc/*any", echoSwagger.WrapHandler)
}

