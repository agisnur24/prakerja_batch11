package main

import (
	"prakerja_batch11/config"
	"prakerja_batch11/route"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	config.InitDatabase()

	e := echo.New()
	route.InitRoutes(e)
	addr := "127.0.0.1:3000"
	e.Start(addr)
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
