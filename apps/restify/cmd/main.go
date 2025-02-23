package main

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"go.opentelemetry.io/otel"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
"github.com/gofiber/fiber/v2/middleware/recover"
	"javifood-restify/config"
	"javifood-restify/internal/infrastructure"
	"javifood-restify/internal/infrastructure/database"
	"javifood-restify/internal/infrastructure/handler"

	"github.com/gofiber/swagger"
	_ "javifood-restify/cmd/docs"
)

var tracer = otel.Tracer("restify")

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	config.NewEnv()
}

//	@title			JaviFood Restify
//	@version		1.0
//	@description	Restify API docs
//	@contact.name	Yuri Barsotti
//	@contact.email	contact@yuribarsotti.tech
//	@host			localhost:3000
//	@basepath		/
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	app := fiber.New()

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

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- app.Listen(":3000")
	}()

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	restaurantsV1 := v1.Group("/restaurants")

	handler.NewRestaurantV1Handler(restaurantsV1)

	app.Get("/", func(c *fiber.Ctx) error {
		_, span := tracer.Start(c.Context(), "hello")
		defer span.End()
		log.WithFields(log.Fields{
			"name":  "Hello Route",
			"route": "/",
		}).Info("Request received")
		return c.SendString("Hello, Fiber!")
	})

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}
	app.Shutdown()
	return
}
