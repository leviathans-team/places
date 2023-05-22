package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	"golang-pkg/internal/auth/usecase"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
)

func UserIdentification(ctx *fiber.Ctx) error {
	//header := ctx.GetRespHeader(authorizationHeader)
	header := ctx.Get(authorizationHeader)
	if header == "" {
		ctx.Status(401)
		log.Print(errors.New("empty auth header"))
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("empty auth header"),
			Timestamp: time.Now(),
		})
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		ctx.Status(401)
		log.Print(errors.New("invalid header"))
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("invalid header"),
			Timestamp: time.Now(),
		})
	}

	headers, err := usecase.ParseToken(headerParts[1])
	if err.Err != nil {
		return ctx.JSON(err)
	}
	ctx.Set("userId", strconv.FormatInt(headers.UserId, 10))
	ctx.Set("IsLandLord", strconv.FormatBool(headers.IsLandLord))
	ctx.Set("IsAdmin", strconv.FormatInt(headers.AdminLevel, 10))

	return nil
}
