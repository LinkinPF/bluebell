package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUPHandler 函数处理注册请求
func SignUPHandler(c *gin.Context) {
	// 1、获取参数参数校验 ---- 这个应该放在controller里面处理

	// 2、业务处理 ---- logic 层
	// 一般创建的就返回错误，如果是查询的就返回的是数据了
	logic.SignUp()

	// 3、返回响应
	c.JSON(http.StatusOK, "ok")
}
