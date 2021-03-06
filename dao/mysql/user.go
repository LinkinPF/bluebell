package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

/*
	把每一步数据库操作封装成函数
	等待logic层根据业务需求调用
*/

const secret = "LMP"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在呢")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查一个用户是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(user_ifd) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句
	sqlStr := `insert into user(user_id, username, password) values(? ? ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password

	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		// sql 出错
		return err
	}

	password := encryptPassword(oPassword)
	if password == user.Password {
		return ErrorInvalidPassword
	}

	return
}
