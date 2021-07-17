package logic

import (
	"web_app/dao/mysql"
	"web_app/model"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(p *model.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.UserName); err != nil {
		zap.L().Error("用户已存在", zap.Error(err))
		return err
	}
	// 2.生成userid

	userID := snowflake.GetID()
	// 构造User实例
	user := &model.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}

	// 3.保存数据库
	return mysql.InsertUser(user)

}

func Login(p *model.ParamLogin) (aToken, rToken string, user *model.User, err error) {
	user = &model.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	//查询用户名密码
	err = mysql.FindUser(user)

	if err != nil {

		return "", "", nil, err
	}
	//生成token
	aToken, rToken, err = jwt.GenToken(user.UserID, user.UserName)
	return

}
