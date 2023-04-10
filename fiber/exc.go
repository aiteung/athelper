package fiber

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {

	response := ReturnData[*string]{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Data:    nil,
		Status:  "Internal Server Error",
	}

	fmt.Printf("%s", err.Error())

	if e, ok := err.(*fiber.Error); ok {
		response.Code = e.Code
		response.Status = e.Message
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		strFinal := "Error terjadi pada field :"
		for _, v := range e {
			strFinal += fmt.Sprintf(" %s", v.Field())
		}
		response.Code = fiber.StatusBadRequest
		response.Status = strFinal

	}

	return response.WriteResponseBody(ctx)
}
