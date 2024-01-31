package middleware

import (
	"context"
	"errors"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/tokenprovider/jwt"
	usermodel "food_delivery/modules/user/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(appctx appctx.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appctx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(common.DbTypeUser)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
