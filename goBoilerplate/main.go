package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	helloWorldController "goBoilerplate/goBoilerplate/helloWorld/web"
)

func main() {
	// fmt.Println(helloWorldController.HelloWorld())
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorldController.HelloWorldHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
