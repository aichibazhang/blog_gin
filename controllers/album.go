package controllers

import (
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func AlbumGet(c *gin.Context) {
	isLogin := c.MustGet("isLogin")
	albums,_:=models.GetAlbums()
	fmt.Println(&albums[0].FilePath)
	c.HTML(http.StatusOK, "album.html", gin.H{"isLogin": isLogin,"albums":albums})
}
func AlbumPost(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		errorResponse(c, err)
		return
	}
	fileType := "otherType"
	fileExt := filepath.Ext(fileHeader.Filename)
	fmt.Println("文件尾缀为:", fileExt)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".bmp" {
		fileType = "img"
	}
	now:=time.Now()
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	fmt.Println(fileDir)
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		errorResponse(c, err)
		return
	}

	fileTime := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", fileTime, fileHeader.Filename)
	filePathStr := fileDir+fileName
	if fileType == "img" {
		album := models.Album{FilePath: filePathStr, FileName: fileName, CreateTime: time.Now().Unix(), Status: 0}
		models.AddAlbum(&album)
	}
	c.SaveUploadedFile(fileHeader, filePathStr)
	c.JSON(http.StatusOK, gin.H{"message": "上传成功", "code": "1"})

}

func errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"response": err})
}
