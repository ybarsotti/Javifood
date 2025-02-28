package handler

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		FailedField string
		Tag         string
		Value       any
	}

	HandlerValidator struct {
		validationErrors map[string]ErrorResponse
	}
)

func NewHandlerValidator() *HandlerValidator {
	return &HandlerValidator{
		validationErrors: make(map[string]ErrorResponse),
	}
}

// TODO: Fix weird messsage
func (h *HandlerValidator) Validate(data any) *HandlerValidator {
	var validate = validator.New()
	errs := validate.Struct(data)

	if errs == nil {
		return h
	}

	for _, err := range errs.(validator.ValidationErrors) {
		elem := new(ErrorResponse)
		elem.Tag = err.Tag()
		elem.Value = err.Value()
		h.validationErrors[err.Field()] = *elem
	}

	return h
}

func (h *HandlerValidator) HasError() bool {
	return len(h.validationErrors) > 0
}

func (h *HandlerValidator) ToFiber() *fiber.Error {
	var errMsgs []string
	for _, err := range h.validationErrors {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}
	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, ","),
	}
}
