package handler

import (
	"javifood-restify/internal/infrastructure/gateway"
	"javifood-restify/internal/interactor"

	fiber "github.com/gofiber/fiber/v2"
)

func NewRestaurantV1Handler(r fiber.Router) {
	restaurantRepository := gateway.NewRestaurantRepository()

	r.Post("", func(c *fiber.Ctx) error {
		h := NewCreateRestaurantHandler(
			interactor.NewCreateRestaurantInteractor(restaurantRepository),
		)
		return h.Handle(c)
	})
}
