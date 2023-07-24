package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Feruz666/advertising/configs"
	"github.com/Feruz666/advertising/pkg/handler"
	"github.com/Feruz666/advertising/pkg/service"
	"github.com/Feruz666/advertising/pkg/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// store
	cfg := configs.Get()

	// Init repository store (with postgresql)
	store, err := store.New(ctx)
	if err != nil {
		return fmt.Errorf("store.New failed %w", err)
	}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return fmt.Errorf("manager.New failed %w", err)
	}

	// Init handlers
	postHandler := handler.NewPost(ctx, serviceManager)

	e := echo.New() 

	e.Use(middleware.Logger())

	postRoutes := e.Group("/adverts")
	postRoutes.GET("/", postHandler.Get)

	s := &http.Server{
		Addr: cfg.ServerAddress,
		ReadTimeout: 30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}


	e.Logger.Fatal(e.StartServer(s))


	return nil
}