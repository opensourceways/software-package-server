package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/opensourceways/software-package-server/common/controller"
	"github.com/opensourceways/software-package-server/softwarepkg/app"
)

type SoftwareController struct {
	controller.BaseController
	repo app.SoftwarePkgService
}

func AddRouteForSoftwareController(r *gin.RouterGroup, repo app.SoftwarePkgService) {
	ctl := SoftwareController{
		repo: repo,
	}

	r.POST("/v1/add-software", ctl.AddSoftware)

}

// AddSoftware
// @Summary add software
// @Description add software
// @Tags  softwarePkg
// @Accept json
// @Param	param  body	 softwareRequest	true	"body of creating software pkg"
// @Success 201 {object} controller.ResponseData
// @Failure 400 {object} controller.ResponseData
// @Router /v1/add-software [post]
func (ctl SoftwareController) AddSoftware(ctx *gin.Context) {
	var req softwareRequest

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctl.SendBadRequestBody(ctx, err)

		return
	}

	pkg, user, err := req.ToCmd()
	if err != nil {
		ctl.SendBadRequestParam(ctx, err)

		return
	}

	if code, cErr := ctl.repo.ApplyNewPkg(user, &pkg); cErr != nil {
		ctl.SendBadRequest(ctx, code, cErr)
	} else {
		ctl.SendCreateSuccess(ctx)
	}
}
