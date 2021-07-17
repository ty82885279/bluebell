package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	//token
	CodeInvalidToken
	CodeTokenInvalidFormat
	CodeNeedLogin
	CodeATokenNotExpir
	CodeRTokenHasExpir
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeInvalidParam:    "参数错误",
	CodeUserExist:       "用户已注册",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务器繁忙",
	//token
	CodeNeedLogin:          "请登陆",
	CodeTokenInvalidFormat: "Token格式错误",
	CodeInvalidToken:       "AToken已过期",
	CodeATokenNotExpir:     "AToken尚未过期",
	CodeRTokenHasExpir:     "RToken已过期",
}

func (r ResCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
