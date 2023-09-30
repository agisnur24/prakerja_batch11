package main

import (
	"os"
	"prakerja_batch11/config"
	"prakerja_batch11/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDatabase()

	e := echo.New()
	route.InitRoutes(e)
	e.Start(":" + os.Getenv("PORT"))
}
