package controllers

import (
	logger "blogweb_gin/gb"
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func AddArticleGet(c *gin.Context) {
	isLogin := GetSession(c)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"isLogin": isLogin})
}

// AddArticlePost 添加文章
func AddArticlePost(c *gin.Context) {
	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	logger.Debug("AddArticlePost", zap.String("title", title), zap.String("tags", tags))

	//实例化model，将它出入到数据库中
	art := &models.Article{
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     "chinadong",
		CreateTime: time.Now().Unix(),
	}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		logger.Error("AddArticlePost failed", zap.Any("error", err))
		response = gin.H{"code": 0, "message": "error"}
	}
	c.JSON(http.StatusOK, response)
}
func GetArticleInfo(c *gin.Context) {
	isLogin := c.MustGet("isLogin")
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("用户获取文章id为:%d\n", id)
	article, _ := models.GetArticleById(id)
	logger.Debug("获取文章详情", zap.Any("article:", article))
	c.HTML(http.StatusOK, "show_article.html",
		gin.H{"isLogin": isLogin, "Title": article.Title, "Content": utils.SwitchMarkdownToHtml(article.Content)})
}
func UpdateArticleGet(c *gin.Context) {
	isLogin := c.MustGet("isLogin")
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Printf("用户获取文章id为:%d\n", id)
	art, _ := models.GetArticleById(id)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": isLogin,
		"article": art})
}
func UpdateArticlePost(c *gin.Context) {
	user := c.MustGet("login").(string)
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("postid:", id)
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	art := &models.Article{
		Author:  user,
		Id:      id,
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
	}
	logger.Debug("UpdateArticlePost", zap.Any("article", *art))
	_, err := models.UpdateArticle(art)
	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "更新成功"}
	} else {
		logger.Error("UpdateArticlePost error", zap.Error(err))
		response = gin.H{"code": 0, "message": "更新失败"}
	}

	c.JSON(http.StatusOK, response)
}
func DeleteArticle(c *gin.Context) {
	idStr := c.Query("id")
	_, err := models.DeleteArticle(idStr)
	if err != nil {
		logger.Error("删除文章失败",zap.Any("error",err))
	}

	c.Redirect(http.StatusFound, "/")
}
