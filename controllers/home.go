package controllers

import (
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeGet(c *gin.Context) {
	isLogin := GetSession(c)
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	var artList []*models.Article
	artList, _ = models.FindArticleWithPage(page)
	fmt.Println("当前用户是否登录:", isLogin)
	data := models.GenHomeBlocks(artList, isLogin)
	homeFooterCode := models.ConfigHomeFooterPageCode(page)
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin, "data": data,"pageData": homeFooterCode})
}
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		return true
	}
	return false
}
