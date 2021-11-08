package middleware

import (
	"fmt"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Guard(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		//db := appCtx.GetMaiDBConnection()
		//store := userstore.NewSQLStore(db)
		//
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		//
		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		// user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		// if err != nil {
		// 	panic(err)
		// }

		c.Set(common.CurrentUser, payload)
		c.Next()
	}
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}
