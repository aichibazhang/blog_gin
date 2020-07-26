package controllers

import (
	logger "blogweb_gin/gb"
	"blogweb_gin/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
		Title:title,
		Tags:tags,
		Short:short,
		Content: content,
		Author:"chinadong",
		CreateTime:time.Now().Unix(),
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

