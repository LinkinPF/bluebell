package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

/*
	存放业务逻辑的代码，会多次调用 dao 层的
*/

func SignUp(p *models.ParamSignUp) (err error) {
	// 1、判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2、生成user_id
	userID := snowflake.GenID()
	// 构造一个 User 实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3、保存进数据库
	return mysql.InsertUser(&user)
}
