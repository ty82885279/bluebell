package controllers

import (
	"net/http"
	"strconv"
	"web_app/logic"
	"web_app/model"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

//CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {

	//校验参数
	var p = new(model.Post)
	err := c.ShouldBind(p)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//业务处理
	userID, err := GetCurrentUserID(c) //获取userID
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("create post failed", zap.Error(err))

		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

//PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	//1. 参数校验
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("PostDetailHandler 参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2. 查询帖子详情
	post, err := logic.GetPostDetailByID(postID)

	if err != nil {
		zap.L().Error("logic.GetPostDetailByID", zap.Int64("id", postID), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3. 返回响应
	ResponseSuccess(c, post)

}

//PostListHandler 帖子列表
func PostListHandler(c *gin.Context) {

	page, size := GetPageInfo(c)

	//查询帖子
	postsData, err := logic.GetPostList(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postsData)

}

// PostListHandler2 根据排序返回帖子列表
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query model.ParamPostList true "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func PostListHandler2(c *gin.Context) {
	http.ServeContent()
	p := &model.ParamPostList{ //默认值
		Page:  1,
		Size:  10,
		Order: model.OrderByTime,
	}
	err := c.ShouldBindQuery(p) //获取请求参数
	if err != nil {

		ResponseError(c, CodeInvalidParam)
		return
	}

	postsData, err := logic.GetPostListNew(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postsData)
}

// CommunityIDPostListHandler 根据社区，按排序返回帖子列表
//func CommunityIDPostListHandler(c *gin.Context) {
//	p := &model.ParamCommunityIDPostList{ //默认值
//		ParamPostList: &model.ParamPostList{
//			Order: model.OrderByTime,
//			Page:  1,
//			Size:  10,
//		},
//	}
//	err := c.ShouldBindQuery(p) //获取请求参数
//	zap.L().Debug("CommunityIDPostListHandler", zap.Any("CommunityID", p.CommunityID), zap.Any("Page", p.Page), zap.Any("Size", p.Size), zap.Any("Order", p.Order))
//	if err != nil {
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	postsData, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, postsData)
//}
