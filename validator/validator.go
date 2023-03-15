package validator

import (
	"fmt"
	playval "github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	goval "github.com/gookit/validate"
)

func ParseAndValidateGooKit[T any](body []byte, data *T) (err error) {
	err = json.Unmarshal(body, data)
	if err != nil {
		return
	}
	v := goval.Struct(data)
	if !v.Validate() {
		return v.Errors.OneError()
	}
	return
}

func ParseAndValidatePlayGround[T any](body []byte, data *T) (err error) {
	err = json.Unmarshal(body, data)
	if err != nil {
		return
	}
	err = playval.New().Struct(data)
	validationErrors := err.(playval.ValidationErrors)
	if validationErrors != nil {
		err = fmt.Errorf("%s", validationErrors.Error())
		return
	}
	return
}
