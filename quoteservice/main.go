package main

import (
	"github.com/dimiro1/health"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	health.HealthAPI(e)
}
