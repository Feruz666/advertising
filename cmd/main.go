package main

import (
	"log"
	"net/http"

	"github.com/Feruz666/advertising/configs"
	"github.com/labstack/echo/v4"
)

func main() {


	config, err := configs.LoadConfig("../")
	if err != nil {
		log.Fatal("cannon load config:", err)
	}

	e := echo.New()
	e.GET("/healthchecker", func(ctx echo.Context) error {
		message := "Welcome to Golang with Postgres"
		return ctx.String(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(config.ServerAddress))

}
