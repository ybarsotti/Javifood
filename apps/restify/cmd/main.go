package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"

	"go.opentelemetry.io/otel"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"

	"javifood-restify/internal/infrastructure"
)

var tracer = otel.Tracer("restify")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	app := fiber.New()
	app.Use(otelfiber.Middleware())

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("ccc")
		return c.SendString("Hello, Fiber!")
	})

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

	select {
	case err = <- srvErr:
		return
	case <-ctx.Done():
		stop()
	}
	app.Shutdown()
	return
}
