package main

import (
	"net/http"

	_ "github.com/alfanherya/pemprograman-golang/golangdocumentation/docs/golangdocumentation"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// add routes
	e.GET("/", HealtCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "hello herya!")
	// })
	e.Logger.Fatal(e.Start(":3000"))
}

func HealtCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "server is up and running",
	})
}
