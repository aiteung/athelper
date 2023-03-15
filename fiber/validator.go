package fiber

import (
	"fmt"
	playval "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	goval "github.com/gookit/validate"
)

func ParseAndValidateGooKit[T any](ctx *fiber.Ctx, data *T) error {
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}

	v := goval.Struct(data)
	if !v.Validate() {
		return v.Errors.OneError()
	}
	return nil
}

func ParseAndValidatePlayGround[T any](ctx *fiber.Ctx, data *T) error {
	err := ctx.BodyParser(data)
	if err != nil {
		return err
	}

	err = playval.New().Struct(data)
	validationErrors := err.(playval.ValidationErrors)

	return fmt.Errorf("%s", validationErrors.Error())
}
