package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/model"
	"web_app/pkg/jwt"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 用户注册
// @Summary 注册
// @Description 用户注册接口
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param object body model.ParamSignUp true "注册参数"
// @Success 200 {object} _ResponseCodeWithMsg "成功数据"
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	var p model.ParamSignUp

	err := c.ShouldBind(&p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	if err = logic.SignUp(&p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 用户登陆
// @Summary 登陆
// @Description 用户登陆接口
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param object body model.ParamLogin true "登陆参数"
// @Success 200 {object} _ResponseLogin "登陆成功"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	//参数校验
	var p model.ParamLogin
	err := c.ShouldBind(&p)
	fmt.Println(p)
	if err != nil {

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	aToken, rToken, user, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("登陆失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		return
	}
	//返回响应
	ResponseSuccess(c, gin.H{
		"msg":      "登陆成功",
		"username": user.UserName,
		"userID":   user.UserID,
		"AToken":   aToken,
		"RToken":   rToken,
	})

}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	rToken := c.Query("refresh_token")
	aToken := c.Request.Header.Get("Authorization")
	if aToken == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求中缺少Auth Token")
		return
	}
	//按空格分割
	parts := strings.SplitN(aToken, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseError(c, CodeTokenInvalidFormat)
		return
	}
	newAToken, newRToken, err := jwt.RefreshToken(parts[1], rToken)
	if err != nil {
		if errors.Is(err, jwt.ErrorRTokenHasExpir) {
			ResponseError(c, CodeRTokenHasExpir)
			return
		}
		if errors.Is(err, jwt.ErrorATokenNotExpir) {
			ResponseError(c, CodeATokenNotExpir)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAToken,
		"refresh_token": newRToken,
	})

}
