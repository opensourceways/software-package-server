package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/community-robot-lib/utils"

	commonstl "github.com/opensourceways/software-package-server/common/controller"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

var (
	userinfoUrl string
	client      utils.HttpClient
)

const (
	headerPrivateToken = "PRIVATE-TOKEN"
	yG                 = "_Y_G_"
	user               = "user"
)

func Init(url string) {
	client = utils.NewHttpClient(3)
	userinfoUrl = url
}

func CheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(headerPrivateToken)
		if len(token) == 0 {
			ctx.JSON(http.StatusBadRequest, commonstl.NewBadRequestHeader("no token"))
			ctx.Abort()

			return
		}

		cookie, err := ctx.Cookie(yG)
		if err != nil || len(cookie) == 0 {
			ctx.JSON(http.StatusBadRequest, commonstl.NewBadRequestCookie("no cookie"))
			ctx.Abort()

			return
		}

		req, _ := http.NewRequest("GET", userinfoUrl, nil)
		req.Header.Set("token", token)
		req.Header.Set("Cookie", yG+"="+cookie)

		var result = struct {
			Data struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"data"`
		}{}

		code, _ := client.ForwardTo(req, &result)
		if code == http.StatusUnauthorized {
			ctx.JSON(http.StatusUnauthorized, commonstl.NewBadRequest("no login"))
			ctx.Abort()

			return
		}

		var userinfo domain.User
		userinfo.Account, err = dp.NewAccount(result.Data.Username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, commonstl.NewBadRequest(err.Error()))
			ctx.Abort()

			return
		}

		userinfo.Email, err = dp.NewEmail(result.Data.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, commonstl.NewBadRequest(err.Error()))
			ctx.Abort()

			return
		}

		ctx.Set(user, &userinfo)

		ctx.Next()
	}
}

func GetUser(ctx *gin.Context) (*domain.User, error) {
	u, _ := ctx.Get(user)
	if userinfo, ok := u.(*domain.User); ok {
		return userinfo, nil
	}

	return nil, errors.New("no userinfo")
}
