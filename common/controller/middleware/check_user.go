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

type Api struct {
	UserInfoURL string `json:"user_info_url" required:"true"`
}

const (
	headerPrivateToken = "PRIVATE-TOKEN"
	yG                 = "_Y_G_"
	user               = "userinfo"
)

func Init(url string) {
	client = utils.NewHttpClient(3)
	userinfoUrl = url
}

func CheckUser(ctx *gin.Context) {
	t, err := token(ctx)
	if err != nil {
		commonstl.SendBadRequestHeader(ctx, err)
		ctx.Abort()

		return
	}

	c, err := cookie(ctx)
	if err != nil {
		commonstl.SendBadRequestCookie(ctx, err)
		ctx.Abort()

		return
	}

	userinfo, err := getUserInfo(t, c)
	if err != nil {
		commonstl.SendBadRequest(ctx, "", err)
		ctx.Abort()

		return
	}

	ctx.Set(user, userinfo)

	ctx.Next()
}

func GetUser(ctx *gin.Context) (*domain.User, error) {
	u, _ := ctx.Get(user)
	if userinfo, ok := u.(*domain.User); ok {
		return userinfo, nil
	}

	return nil, errors.New("no userinfo")
}

func token(ctx *gin.Context) (t string, err error) {
	if t = ctx.GetHeader(headerPrivateToken); len(t) == 0 {
		err = errors.New("invalid token")
	}

	return
}

func cookie(ctx *gin.Context) (c string, err error) {
	if c, err = ctx.Cookie(yG); err != nil || len(c) == 0 {
		err = errors.New("invalid cookie")
	}

	return
}

func getUserInfo(t, c string) (userinfo *domain.User, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", userinfoUrl, nil)
	if err != nil {
		return
	}

	req.Header.Set("token", t)
	req.Header.Set("Cookie", yG+"="+c)

	var result = struct {
		Data struct {
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"data"`
	}{}

	code, _ := client.ForwardTo(req, &result)
	if code == http.StatusUnauthorized {
		err = errors.New("no login")

		return
	}

	userinfo = new(domain.User)
	userinfo.Account, err = dp.NewAccount(result.Data.Username)
	if err != nil {
		return
	}

	userinfo.Email, err = dp.NewEmail(result.Data.Email)
	if err != nil {
		return
	}

	return
}
