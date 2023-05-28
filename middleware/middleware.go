package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	"golang-pkg/internal/auth/usecase"
	userRepostiory "golang-pkg/internal/user/repository"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
)

// UserIdentification
// @Summary Проверка токена
// @Tags auth
// @Description На вход получаю токен в хедере Authorization
// @ID UserIdentification
// @Param        Authorization   header      string  true  "bearer token"
// @Produce  json
// @Success 200 {object} string
// @Failure 500 {object} internal.HackError
// @Router /getUserInfo [get]
// Возвращаю jwt-token, который храним в себе данные об авторизации
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
	if err != nil {
		return ctx.JSON(err)
	}
	ctx.Set("userId", strconv.FormatInt(headers.UserId, 10))
	ctx.Set("IsLandLord", strconv.FormatBool(headers.IsLandLord))
	ctx.Set("adminLevel", strconv.FormatInt(headers.AdminLevel, 10))

	return ctx.Next()
}

func UserIsExist(ctx *fiber.Ctx) error {
	userId := ctx.GetRespHeader("userId", "")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	res, hackErr := userRepostiory.IsExistsOnUsersTable(userIdInt)
	if hackErr != nil {
		return ctx.JSON(hackErr)
	}
	if !res {
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("user not found"),
			Message:   "user not found",
			Timestamp: time.Now(),
		})
	}
	return ctx.Next()
}

func AdminIsExist(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}

	res, hackErr := userRepostiory.IsExistsOnAdminTable(adminIdInt)
	if hackErr != nil {
		return ctx.JSON(hackErr)
	}
	if !res {
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("user not found"),
			Message:   "user not found",
			Timestamp: time.Now(),
		})
	}
	return ctx.Next()
}

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(user interface{}) []*ErrorResponse {
	var errorResponses []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errorResponses = append(errorResponses, &element)
		}
	}
	return errorResponses
}
