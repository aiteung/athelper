package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {

	response := ReturnData[*string]{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Data:    nil,
	}

	fmt.Printf("%s", err.Error())

	if e, ok := err.(*fiber.Error); ok {
		response.Code = e.Code
		response.Status = e.Message
	}

	return response.WriteResponseBody(ctx)
}
