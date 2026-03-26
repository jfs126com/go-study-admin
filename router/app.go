package router

import (
	"github.com/gin-gonic/gin"
	"go-study-admin/service"
)

func App() *gin.Engine {
	r := gin.Default()
	// 注册登录路由
	r.POST("/api/login/password", service.LoginPassword)
	return r
}
