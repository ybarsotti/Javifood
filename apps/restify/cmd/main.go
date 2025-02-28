package main

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"javifood-restify/config"
	"javifood-restify/internal/infrastructure"
	"javifood-restify/internal/infrastructure/database"
	"javifood-restify/internal/infrastructure/handler"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"

	_ "javifood-restify/cmd/docs"

	"github.com/gofiber/swagger"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	config.NewEnv()
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// @title			JaviFood Restify
// @version			1.0
// @description		Restify API docs
// @contact.name	Yuri Barsotti
// @contact.email	contact@yuribarsotti.tech
// @host			localhost:3000
// @basepath		/
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(otelfiber.Middleware())
	app.Use(cors.New())
	app.Use(recover.New())

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := infrastructure.SetupOTelSDK(ctx)
	if err != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	err = database.InitDatabase()

	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking: true,
		Title:       "Restify API Docs",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	restaurantsV1 := v1.Group("/restaurants")

	handler.NewRestaurantV1Handler(restaurantsV1)

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- app.Listen(":3000")
	}()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}
	err = app.Shutdown()
	log.Error(err)
	return
}
