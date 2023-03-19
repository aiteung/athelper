package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

type AuthMiddleware struct {
	PublicKey        string
	PrivateKey       string
	AuthHeader       string
	AuthHeaderDecode string
}

func InitAuthMiddleware(publicKey string, privateKey string, authHeader string, authHeaderDecode string) *AuthMiddleware {
	return &AuthMiddleware{
		PublicKey:        publicKey,
		PrivateKey:       privateKey,
		AuthHeader:       authHeader,
		AuthHeaderDecode: authHeaderDecode,
	}
}

func (auth *AuthMiddleware) DecodeToken(ctx *fiber.Ctx) (err error) {
	tokenString := ctx.Get(auth.AuthHeader)

	if tokenString == "" {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Missing %s Header", auth.AuthHeader))
	}

	payload, err := watoken.Decode(auth.PublicKey, tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorize Token")
	}

	ctx.Set(auth.AuthHeaderDecode, payload.Id)
	err = ctx.Next()
	return
}