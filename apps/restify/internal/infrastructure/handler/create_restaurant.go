package handler

import (
	"javifood-restify/internal/domain/usecase"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var t = otel.Tracer("restaurant")

type (
	CreateRestaurantHandler struct {
		usecase usecase.CreateRestaurantUseCase
	}
	payloadDto struct {
		UserID      string   `json:"user_id"      example:"01953aa5-7d18-7781-bf1c-f425606b565f"`
		Name        string   `json:"name"         example:"Restaurant XYZ"                       validate:"required"`
		Address     string   `json:"address"      example:"3598 Pringle Drive"                   validate:"required"`
		CoordinateX float64  `json:"coordinate_x" example:"-21.74568"                            validate:"required"`
		CoordinateY float64  `json:"coordinate_y" example:"-89.34886"                            validate:"required"`
		OpenTime    string   `json:"open_time"    example:"10:00"                                validate:"required"`
		CloseTime   string   `json:"close_time"   example:"22:30"                                validate:"required"`
		WorkDays    []string `json:"work_days"    example:"['Monday', 'Tuesday']"                validate:"required" enums:"Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"`
	}
)

// CreateRestaurant godoc
//
//	@Summary		Create a restaurant
//	@Description	Allows to create a restaurant to the user
//	@Tags			restaurants
//	@Accept			json
//	@Success		201
//	@Router			/api/v1/restaurants/ [post]
//	@Param			x-user header string true "User ID from Clerk JWT"
//	@Param			data	body	payloadDto	true	"Restaurant data"
func (h *CreateRestaurantHandler) Handle(c *fiber.Ctx) error {
	_, span := t.Start(c.Context(), "create_restaurant")
	defer span.End()
	payload := &payloadDto{}
	if err := c.BodyParser(&payload); err != nil {
		log.Error("Failed parsing request body: ", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.UserID = c.Get("x-user")
	validatedData := NewHandlerValidator().Validate(payload)
	if err := validatedData.HasError(); err {
		log.Error(validatedData.ToFiber().Error())
		return validatedData.ToFiber()
	}
	err := h.usecase.Execute(c.Context(), *usecase.NewCreateRestaurantInputDto(
		payload.UserID, payload.Name, payload.Address, payload.OpenTime, payload.CloseTime, payload.CoordinateX, payload.CoordinateY, payload.WorkDays,
	))
	if err != nil {
		// TODO: Add specific error handling
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(201)
}

func NewCreateRestaurantHandler(usecase usecase.CreateRestaurantUseCase) *CreateRestaurantHandler {
	return &CreateRestaurantHandler{
		usecase: usecase,
	}
}
