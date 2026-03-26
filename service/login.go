package service

import (
	"go-study-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPassword(c *gin.Context) {
	in := new(LoginPasswordRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	}
	sysUser, err := models.GetUserByUsernamePass(in.Username, in.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	// 登录成功，返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": sysUser,
	})
}
