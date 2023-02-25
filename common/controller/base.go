package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BaseController struct {
}

func (ctl BaseController) SendBadRequestBody(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, newResponseCodeMsg(errorBadRequestBody, err.Error()))
}

func (ctl BaseController) SendBadRequestParam(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, newResponseCodeMsg(errorBadRequestParam, err.Error()))
}

func (ctl BaseController) SendCreateSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, newResponseCodeMsg("", "success"))
}

func (ctl BaseController) SendBadRequest(ctx *gin.Context, code string, err error) {
	if code == "" {
		ctx.JSON(http.StatusBadRequest, newResponseCodeMsg(errorBadRequest, err.Error()))

		return
	}

	ctx.JSON(http.StatusBadRequest, newResponseCodeMsg(code, err.Error()))
}

func (ctl BaseController) SendBadRequestParamWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, newResponseCodeMsg(errorBadRequestParam, msg))
}

func (ctl BaseController) SendRespOfGet(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResponseData(data))
}

func (ctl BaseController) SendRespOfPost(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, newResponseData(data))
}

func (ctl BaseController) SendRespOfPut(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusAccepted, newResponseData(data))
}

func (ctl BaseController) SendRespOfDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, newResponseData("success"))
}

func (ctl BaseController) SendRespWithInternalError(ctx *gin.Context, data ResponseData) {
	logrus.Errorf("code: %s, err: %s", data.Code, data.Msg)

	ctx.JSON(http.StatusInternalServerError, data)
}
