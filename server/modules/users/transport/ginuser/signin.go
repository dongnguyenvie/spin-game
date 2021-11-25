package ginuser

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	"nolan/spin-game/components/tokenprovider/jwt"
	userbiz "nolan/spin-game/modules/users/biz"
	usermodel "nolan/spin-game/modules/users/model"
	userrepo "nolan/spin-game/modules/users/repository"
	userstorage "nolan/spin-game/modules/users/storage"
	"time"

	"github.com/gin-gonic/gin"
)

func SignIn(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.GetUserByWallet

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMaiDBConnection()
		userStorage := userstorage.NewSQLStore(db)
		signinRepo := userrepo.NewUserRepo(userStorage)
		signinBiz := userbiz.NewSigninBiz(signinRepo)

		user, err := signinBiz.Signin(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		secretKey := appCtx.GetSecretKey()

		jwtProvider := jwt.NewTokenJWTProvider(secretKey)

		token, err := jwtProvider.Generate(tokenprovider.TokenPayload{UserId: user.Id, WalletAddress: user.WalletAddress, Email: user.Email}, int(time.Minute*20))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))

	}
}
