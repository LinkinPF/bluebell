package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// SignUPHandler 函数处理注册请求
func SignUPHandler(c *gin.Context) {
	//  1、获取参数和参数校验 ---- 这个应该放在 controller 里面处理
	//	前后端分离，是JSON格式的数据，所以使用 c.ShouldBind() 方法;  如果不是前后端分离，就使用 c.Query()  c.Param()
	//	然后就需要定义请求参数的结构体，在models文件夹下定义一个 params.go，专门用来定义参数结构体。
	var p models.ParamSignUp
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"message": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	// 2、业务处理 ---- logic 层
	// 一般创建的就返回错误，如果是 查询的就返回的是数据了
	logic.SignUp()

	// 3、返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
