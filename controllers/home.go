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
	var homeFooterCode models.HomeFooterPage
	tag := c.Query("tag")
	if len(tag) > 0 {
		artList, _ = models.QueryArticlesWithTag(page, tag)
	} else {
		artList, _ = models.FindArticleWithPage(page)
	}
	homeFooterCode = models.ConfigHomeFooterPageCode(page)
	fmt.Println("当前用户是否登录:", isLogin)
	data := models.GenHomeBlocks(artList, isLogin)
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin, "data": data, "pageData": homeFooterCode})
}
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		return true
	}
	return false
}
