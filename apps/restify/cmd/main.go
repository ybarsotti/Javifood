package main

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"go.opentelemetry.io/otel"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"

	"javifood-restify/config"
	"javifood-restify/internal/infrastructure"
)

var tracer = otel.Tracer("restify")

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	config.NewEnv()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	app := fiber.New()

	app.Use(otelfiber.Middleware())

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

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- app.Listen(":3000")
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		_, span := tracer.Start(c.Context(), "hello")
		defer span.End()
		log.WithFields(log.Fields{
			"name": "Hello Route",
			"route": "/",
		}).Info("Request received")
		return c.SendString("Hello, Fiber!")
	})

	select {
	case err = <- srvErr:
		return
	case <-ctx.Done():
		stop()
	}
	app.Shutdown()
	return
}
