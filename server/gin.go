package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/community-robot-lib/interrupts"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/opensourceways/software-package-server/config"
	"github.com/opensourceways/software-package-server/docs"
	"github.com/opensourceways/software-package-server/infrastructure/postgresql"
	softwarepkgapp "github.com/opensourceways/software-package-server/softwarepkg/app"
	"github.com/opensourceways/software-package-server/softwarepkg/controller"
	"github.com/opensourceways/software-package-server/softwarepkg/infrastructure/repositoryimpl"
)

func StartWebServer(port int, timeout time.Duration, cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logRequest())

	setRouter(r, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	defer interrupts.WaitForGracefulShutdown()

	interrupts.ListenAndServe(srv, timeout)
}

//setRouter init router
func setRouter(engine *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "Software Package"
	docs.SwaggerInfo.Description = "set header: 'PRIVATE-TOKEN=xxx'"

	v1 := engine.Group(docs.SwaggerInfo.BasePath)
	setApiV1(v1)

	engine.UseRawPath = true
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func setApiV1(v1 *gin.RouterGroup) {
	initSoftwarePkgService(v1)
}

func initSoftwarePkgService(v1 *gin.RouterGroup) {
	controller.AddRouteForSoftwareController(v1, checkUser(), softwarepkgapp.NewSoftwarePkgService(
		repositoryimpl.NewSoftware(postgresql.DB())))

}

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		logrus.Infof(
			"| %d | %d | %s | %s |",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}

func checkUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
		ut, _ := c.Cookie("_U_T_")
		yg, _ := c.Cookie("_Y_G_")

		if len(ut) == 0 || len(yg) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "bad_request",
				"msg":  "invalid cookie",
			})

			c.Abort()

			return
		}

		var (
			resp *http.Response
			err  error
		)

		// TODO to be provided
		resp, err = http.Post("", "application/json", nil)
		if err == nil && resp.StatusCode == http.StatusOK {
			c.Next()
		} else {
			if err != nil {
				logrus.Errorf("check user faild,err :%v", err)
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "bad_request",
				"msg":  "invalid user",
			})

			c.Abort()

			return
		}
	}
}
