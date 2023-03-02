package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	commonctl "github.com/opensourceways/software-package-server/common/controller"
	"github.com/opensourceways/software-package-server/softwarepkg/app"
)

type SoftwarePkgController struct {
	commonctl.BaseController
	service app.SoftwarePkgService
}

func AddRouteForSoftwarePkgController(r *gin.RouterGroup, pkgService app.SoftwarePkgService) {
	ctl := SoftwarePkgController{
		service: pkgService,
	}

	r.POST("/v1/softwarepkg", ctl.ApplyNewPkg)
	r.GET("/v1/softwarepkg", ctl.ListPkgs)
	r.GET("/v1/:pid/softwarepkg", ctl.Get)
}

// ApplyNewPkg
// @Summary apply a new software package
// @Description apply a new software package
// @Tags  SoftwarePkg
// @Accept json
// @Param	param  body	 softwarePkgRequest	 true	"body of applying a new software package"
// @Success 201 {object} ResponseData
// @Failure 400 {object} ResponseData
// @Router /v1/softwarepkg [post]
func (ctl SoftwarePkgController) ApplyNewPkg(ctx *gin.Context) {
	var req softwarePkgRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctl.SendBadRequestBody(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		ctl.SendBadRequestParam(ctx, err)

		return
	}

	if code, err := ctl.service.ApplyNewPkg(nil, &cmd); err != nil {
		ctl.SendBadRequest(ctx, code, err)
	} else {
		ctl.SendCreateSuccess(ctx)
	}
}

// ListPkgs
// @Summary list software packages
// @Description list software packages
// @Tags  SoftwarePkg
// @Accept json
// @Param    importer         query	 string   false    "importer of the softwarePkg"
// @Param    phase            query	 string   false    "phase of the softwarePkg"
// @Param    count_per_page   query	 int      false    "count per page"
// @Param    page_num         query	 int      false    "page num which starts from 1"
// @Success 200 {object} app.SoftwarePkgsDTO
// @Failure 400 {object} ResponseData
// @Router /v1/softwarepkg [get]
func (ctl SoftwarePkgController) ListPkgs(ctx *gin.Context) {
	var req softwarePkgListQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctl.SendBadRequestParam(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		ctl.SendBadRequestParam(ctx, err)

		return
	}

	if v, err := ctl.service.ListPkgs(&cmd); err != nil {
		ctl.SendBadRequest(ctx, "", err)
	} else {
		ctl.SendRespOfGet(ctx, v)
	}
}

// Get
// @Summary get software package
// @Description get software package
// @Tags  SoftwarePkg
// @Accept json
// @Param    pid         path	string  true    "pid of software package id"
// @Success 200 {object} app.SoftwarePkgReviewDTO
// @Failure 400 {object} ResponseData
// @Router /v1/:pid/softwarepkg [get]
func (ctl SoftwarePkgController) Get(ctx *gin.Context) {
	if v, err := ctl.service.GetPkgReviewDetail(ctx.Param("pid")); err != nil {
		ctl.SendBadRequest(ctx, "", err)
	} else {
		ctl.SendRespOfGet(ctx, v)
	}
}
