package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitGet(c *gin.Context) {
	session:=sessions.Default(c)
	fmt.Println("delete session",session.Get("user"))
	//session.Clear()
	session.Save()
	c.Redirect(http.StatusMovedPermanently,"/")
}
