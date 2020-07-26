package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUser := session.Get("user")
		if loginUser == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("isLogin", true)
		c.Set("login", loginUser)
		c.Next()
	}
}
