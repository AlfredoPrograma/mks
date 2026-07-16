package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alfredoprograma/mks/internal/config"
	"github.com/alfredoprograma/mks/internal/database"
	"github.com/alfredoprograma/mks/internal/users"
	"github.com/labstack/echo/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	queries, err := database.Connect(context.Background(), cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	app := echo.New()
	apiRouter := app.Group("/api/v1")

	usersModule := users.NewModule(queries)
	usersModule.Controller.RegisterRoutes(apiRouter.Group("/users"))

	if err := app.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalln(err)
	}
}
