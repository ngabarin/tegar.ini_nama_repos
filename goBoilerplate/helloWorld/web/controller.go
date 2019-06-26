package helloWorld

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, roadrunner World!")
}
