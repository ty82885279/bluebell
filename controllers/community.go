package controllers

import (
	"strconv"
	"web_app/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 获取所有社区
// @Summary 获取所有社区
// @Description 获取所有社区
// @Tags 社区相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityList
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	//不需要参数处理
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// CommunityDetailHandler 社区详情
// @Summary 社区详情
// @Description 社区详情接口
// @Tags 社区相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int false "社区id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityDetail
// @Router /community/{id} [get]
func CommunityDetailHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		zap.L().Error("communityDetailParmaErr", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.CommunityDetail(id)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	ResponseSuccess(c, data)

}
