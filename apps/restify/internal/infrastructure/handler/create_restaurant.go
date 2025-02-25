package handler

import (
	"fmt"
	"javifood-restify/internal/domain/usecase"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

var t = otel.Tracer("restaurant")

type (
	CreateRestaurantHandler struct {
		usecase usecase.CreateRestaurantUseCase
	}
	payloadDto struct {
		UserID      string   `json:"user_id"      example:"01953aa5-7d18-7781-bf1c-f425606b565f" validate:"required"`
		Name        string   `json:"name"         example:"Restaurant XYZ"                       validate:"required"`
		Address     string   `json:"address"      example:"3598 Pringle Drive"                   validate:"required"`
		CoordinateX float64  `json:"coordinate_x" example:"-21.74568"                            validate:"required"`
		CoordinateY float64  `json:"coordinate_y" example:"-89.34886"                            validate:"required"`
		OpenTime    string   `json:"open_time"    example:"10:00"                                validate:"required"`
		CloseTime   string   `json:"close_time"   example:"22:30"                                validate:"required"`
		WorkDays    []string `json:"work_days"    example:"['Monday', 'Tuesday']"                validate:"required" enums:"Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"`
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
// @Description  Allows to create a restaurant to the user
// @Tags         restaurants
// @Accept       json
// @Success      201
// @Router       /api/v1/restaurants/ [post]
// @Param		 data body payloadDto true "Restaurant data"
func (h *CreateRestaurantHandler) Handle(c *fiber.Ctx) error {
	_, span := t.Start(c.Context(), "create_restaurant")
	defer span.End()
	validate := validator.New()
	payload := &payloadDto{}
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
			errMsgs = append(
				errMsgs,
				fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					elem.FailedField,
					elem.Value,
					elem.Tag,
				),
			)
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
