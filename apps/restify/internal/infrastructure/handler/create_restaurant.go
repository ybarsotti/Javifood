package handler

import (
	"fmt"
	"javifood-restify/internal/domain/usecase"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

var t = otel.Tracer("create_restaurant")

type (
	CreateRestaurantHandler struct {
		usecase usecase.CreateRestaurantUseCase
	}
	payloadDto struct {
		UserID      string   `json:"user_id" validate:"required"`
		Name        string   `json:"name" validate:"required"`
		Address     string   `json:"address" validate:"required"`
		CoordinateX float64  `json:"coordinate_x" validate:"required"`
		CoordinateY float64  `json:"coordinate_y" validate:"required"`
		OpenTime    string   `json:"open_time" validate:"required"`
		CloseTime   string   `json:"close_time" validate:"required"`
		WorkDays    []string `json:"work_days" validate:"required"`
	}
	ErrorResponse struct {
		FailedField string
		Tag         string
		Value       interface{}
	}
)

// CreateRestaurant godoc
//
// @Summary      Create a restaurant
// @Tags         restaurants
// @Accept       json
// @Success      201  
// @Router       /api/v1/restaurants/ [post]
func (h *CreateRestaurantHandler) Handle(c *fiber.Ctx) error {
	_, span := t.Start(c.Context(), "create_restaurant")
	payload := &payloadDto{}
	defer span.End()
	validate := validator.New()
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	errs := validate.Struct(payload)
	if errs != nil {
		errMsgs := make([]string, len(errs.(validator.ValidationErrors)))
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			errMsgs = append(errMsgs, fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", elem.FailedField, elem.Value, elem.Tag))
		}
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, ""),
		}
	}
	// h.usecase.Execute(c.Context(), *usecase.NewCreateRestaurantInputDto(
	// 	payload.UserID, payload.Name, payload.Address, strconv.Atoi(strings.Join([]string{
	// 		payload.CoordinateX
	// 	}))
	// 	))
	return c.SendStatus(201)
}

func NewCreateRestaurantHandler(usecase usecase.CreateRestaurantUseCase) *CreateRestaurantHandler {
	return &CreateRestaurantHandler{
		usecase: usecase,
	}
}
