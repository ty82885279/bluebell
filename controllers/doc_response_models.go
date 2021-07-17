package controllers

import "web_app/model"

type _ResponsePostList struct {
	Code    ResCode                `json:"code"`    // 业务响应状态码
	Message string                 `json:"message"` // 提示信息
	Data    []*model.ApiPostDetail `json:"data"`    // 数据
}
type _ResponseCommunityList struct {
	Code    ResCode            `json:"code"`    // 业务响应状态码
	Message string             `json:"message"` // 提示信息
	Data    []*model.Community `json:"data"`    // 数据
}
type _ResponseCommunityDetail struct {
	Code    ResCode               `json:"code"`    // 业务响应状态码
	Message string                `json:"message"` // 提示信息
	Data    model.CommunityDetail `json:"data"`    // 数据
}

type _ResponseCodeWithMsg struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
}

type _ResponseLogin struct {
	Code     ResCode `json:"code"` // 业务响应状态码
	Message  string  `json:"msg"`  // 提示信息
	UserName string  `json:"user_name"`
	AToken   string  `json:"AToken"`
	RToken   string  `json:"RToken"`
	UseID    int64   `json:"use_id"`
}
