package controllers

import (
	logger "blogweb_gin/gb"
	"blogweb_gin/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetTags(c *gin.Context) {
	tags, err := models.QueryArticleByParam("tags")
	if err != nil {
		logger.Error("获取文章标签失败", zap.Any("error", err))
	}
	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": models.GetTags(tags)})
}
