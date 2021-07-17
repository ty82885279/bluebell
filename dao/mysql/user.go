package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/model"
)

const secret = "lilei"

func CheckUserExist(UserName string) (err error) {
	//select count(user_id) from user where username = ?
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err = DB.Get(&count, sqlStr, UserName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 注册用户
func InsertUser(user *model.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = DB.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return

}

//  FindUser 用户登陆
func FindUser(user *model.User) (err error) {

	opassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = DB.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return errors.New("查询出错")
	}
	//判断密码是否正确
	password := encryptPassword(opassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

//  encryptPassword 密码加盐
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))

}

// GetUserByID 查询用户
func GetUserByID(id int64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = DB.Get(user, sqlStr, id)

	return
}
