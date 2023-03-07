package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/community-robot-lib/utils"

	commonstl "github.com/opensourceways/software-package-server/common/controller"
)

var (
	userinfoUrl string
	client      utils.HttpClient
)

const (
	headerPrivateToken = "PRIVATE-TOKEN"
	yG                 = "_Y_G_"
	giteeUser          = "giteeUser"
	user               = "user"
	email              = "email"
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

		var result Userinfo
		code, _ := client.ForwardTo(req, &result)
		if code == http.StatusUnauthorized {
			ctx.JSON(http.StatusUnauthorized, commonstl.NewBadRequest("no login"))

			ctx.Abort()

			return
		}

		var giteeUserName string
		username := result.Data.Username
		loginEmail := result.Data.Email
		for _, v := range result.Data.Identities {
			if v.Identity == "gitee" {
				giteeUserName = v.LoginName
			}
		}

		ctx.Set(user, username)
		ctx.Set(giteeUser, giteeUserName)
		ctx.Set(email, loginEmail)

		ctx.Next()
	}
}

func UserName(ctx *gin.Context) string {
	return ctx.GetString(user)
}

func GiteeUserName(ctx *gin.Context) string {
	return ctx.GetString(giteeUser)
}

func GetEmail(ctx *gin.Context) string {
	return ctx.GetString(email)
}
