package fiber

import (
	"fmt"
	at "github.com/aiteung/athelper"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

type AuthMiddleware struct {
	PublicKey        string
	PrivateKey       string
	AuthHeader       string
	AuthHeaderDecode string
	Salt             string
}

func NewAuthMiddleware(
	publicKey string,
	privateKey string,
	authHeader string,
	authHeaderDecode string,
	salt string) *AuthMiddleware {
	if authHeader == "" {
		authHeader = "Login"
	}
	if authHeaderDecode == "" {
		authHeader = "Token"
	}

	if salt == "" {
		salt = "+"
	}

	return &AuthMiddleware{
		PublicKey:        publicKey,
		PrivateKey:       privateKey,
		AuthHeader:       authHeader,
		AuthHeaderDecode: authHeaderDecode,
		Salt:             salt,
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

	ctx.Locals(auth.AuthHeaderDecode, at.AddObsToken(payload.Id, auth.Salt))
	err = ctx.Next()
	return
}
