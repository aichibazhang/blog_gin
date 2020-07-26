package controllers

import (
	logger "blogweb_gin/gb"
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"github.com/gin-contrib/sessions"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}
func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	id := models.QueryUserWithParam(username, utils.MD5(password))
	if id > 0 {
		session:=sessions.Default(c)
		session.Set("user",username)
		session.Save()
		logger.Debug("login success",zap.String("login success",username))
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登陆成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登陆失败"})
	}
}
