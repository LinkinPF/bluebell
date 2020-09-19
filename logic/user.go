package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

/*
	存放业务逻辑的代码，会多次调用 dao 层的
*/

func SignUp() {
	// 1、判断用户存不存在
	mysql.QueryUserByUsername()
	// 2、生成user_id
	snowflake.GenID()
	// 3、密码加密

	// 4、保存进数据库
	mysql.InsertUser()
}
