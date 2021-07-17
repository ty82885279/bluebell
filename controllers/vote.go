package controllers

import (
	"web_app/logic"
	"web_app/model"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 投票接口
func PostVoteHandler(c *gin.Context) {

	//获取参数
	vote := new(model.ParamVoteDatas)
	err := c.ShouldBind(vote)
	if err != nil {

		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取当前用户id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	//投票具体业务逻辑
	err = logic.VoteForPost(userID, vote)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}
