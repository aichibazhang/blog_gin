package controllers

import (
	logger "blogweb_gin/gb"
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}
func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "两次密码输入不一致"})
	}
	id := models.QueryUserWithUsername(username)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "哟用户名已存在"})
		return
	}
	password = utils.MD5(password)
	logger.Debug("密码加密",zap.String("md5",password))
	user := models.User{Username: username, Password: password, CreateTime: time.Now().Unix(), Status: 0}
	_, err := models.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}
