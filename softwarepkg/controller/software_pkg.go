package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	commonctl "github.com/opensourceways/software-package-server/common/controller"
	"github.com/opensourceways/software-package-server/common/controller/middleware"
	"github.com/opensourceways/software-package-server/softwarepkg/app"
)

type SoftwarePkgController struct {
	service app.SoftwarePkgService
}

func AddRouteForSoftwarePkgController(r *gin.RouterGroup, pkgService app.SoftwarePkgService) {
	ctl := SoftwarePkgController{
		service: pkgService,
	}

	r.POST("/v1/softwarepkg", middleware.UserChecking().CheckUser, ctl.ApplyNewPkg)
	r.GET("/v1/softwarepkg", ctl.ListPkgs)
	r.GET("/v1/softwarepkg/:id", ctl.Get)

	r.GET("/v1/reviewcomment/approve/:id", middleware.UserChecking().CheckUser, ctl.Approve)
	r.GET("/v1/reviewcomment/reject/:id", middleware.UserChecking().CheckUser, ctl.Reject)
	r.GET("/v1/reviewcomment/abandon/:id", middleware.UserChecking().CheckUser, ctl.Abandon)
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
		commonctl.SendBadRequestBody(ctx, err)

		return
	}

	user, err := middleware.UserChecking().FetchUser(ctx)
	if err != nil {
		commonctl.SendFailedResp(ctx, "", err)

		return
	}

	cmd, err := req.toCmd(&user)
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if code, err := ctl.service.ApplyNewPkg(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, code, err)
	} else {
		commonctl.SendRespOfCreate(ctx)
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
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		commonctl.SendBadRequestParam(ctx, err)

		return
	}

	if v, err := ctl.service.ListPkgs(&cmd); err != nil {
		commonctl.SendFailedResp(ctx, "", err)
	} else {
		commonctl.SendRespOfGet(ctx, v)
	}
}

// Get
// @Summary get software package
// @Description get software package
// @Tags  SoftwarePkg
// @Accept json
// @Param    id         path	string  true    "id of software package"
// @Success 200 {object} app.SoftwarePkgReviewDTO
// @Failure 400 {object} ResponseData
// @Router /v1/softwarepkg/{id} [get]
func (ctl SoftwarePkgController) Get(ctx *gin.Context) {
	if v, err := ctl.service.GetPkgReviewDetail(ctx.Param("id")); err != nil {
		commonctl.SendFailedResp(ctx, "", err)
	} else {
		commonctl.SendRespOfGet(ctx, v)
	}
}

// Approve
// @Summary approve software package
// @Description approve software package
// @Tags  SoftwarePkg
// @Accept json
// @Param	id  path	 string	 true	"id of software package"
// @Success 201 {object} ResponseData
// @Failure 400 {object} ResponseData
// @Router /v1/reviewcomment/approve/{id} [get]
func (ctl SoftwarePkgController) Approve(ctx *gin.Context) {
	user, err := middleware.UserChecking().FetchUser(ctx)
	if err != nil {
		commonctl.SendBadRequest(ctx, "", err)

		return
	}

	if code, err := ctl.service.Approve(ctx.Param("id"), user.Account); err != nil {
		commonctl.SendBadRequest(ctx, code, err)
	} else {
		commonctl.SendCreateSuccess(ctx)
	}
}

// Reject
// @Summary reject software package
// @Description reject software package
// @Tags  SoftwarePkg
// @Accept json
// @Param	id  path	 string	 true	"id of software package"
// @Success 201 {object} ResponseData
// @Failure 400 {object} ResponseData
// @Router /v1/reviewcomment/reject/{id} [get]
func (ctl SoftwarePkgController) Reject(ctx *gin.Context) {
	user, err := middleware.UserChecking().FetchUser(ctx)
	if err != nil {
		commonctl.SendBadRequest(ctx, "", err)

		return
	}

	if code, err := ctl.service.Reject(ctx.Param("id"), user.Account); err != nil {
		commonctl.SendBadRequest(ctx, code, err)
	} else {
		commonctl.SendCreateSuccess(ctx)
	}
}

// Abandon
// @Summary abandon software package
// @Description abandon software package
// @Tags  SoftwarePkg
// @Accept json
// @Param	id  path	 string	 true	"id of software package"
// @Success 201 {object} ResponseData
// @Failure 400 {object} ResponseData
// @Router /v1/reviewcomment/abandon/{id} [get]
func (ctl SoftwarePkgController) Abandon(ctx *gin.Context) {
	user, err := middleware.UserChecking().FetchUser(ctx)
	if err != nil {
		commonctl.SendBadRequest(ctx, "", err)

		return
	}

	if code, err := ctl.service.Abandon(ctx.Param("id"), user.Account); err != nil {
		commonctl.SendBadRequest(ctx, code, err)
	} else {
		commonctl.SendCreateSuccess(ctx)
	}
}
