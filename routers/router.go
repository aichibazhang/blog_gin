package routers

import (
	"blogweb_gin/controllers"
	logger "blogweb_gin/gb"
	"blogweb_gin/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	r.SetFuncMap(template.FuncMap{
		"timeStr": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "static")
	//store := cookie.NewStore([]byte("user"))
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("user"))
	router.Use(sessions.Sessions("MySessions", store))
	{
		router.GET("/register", controllers.Register)
		router.POST("/register", controllers.RegisterPost)
		router.GET("/login", controllers.Login)
		router.POST("/login", controllers.LoginPost)
		router.GET("/exit", controllers.ExitGet)
		router.GET("/top/:n",controllers.TopGet)
	}
	{
		oauthGroup := router.Group("/", middlewares.BasicAuth())
		oauthGroup.GET("/", controllers.HomeGet)
		oauthGroup.GET("/album",controllers.AlbumGet)
		oauthGroup.POST("/upload",controllers.AlbumPost)
		v1 := oauthGroup.Group("/article")
		{
			v1.GET("/add", controllers.AddArticleGet)
			v1.POST("/add", controllers.AddArticlePost)
			v1.GET("/show/:id", controllers.GetArticleInfo)
			v1.GET("/update", controllers.UpdateArticleGet)
			v1.POST("/update", controllers.UpdateArticlePost)
			v1.GET("/delete",controllers.DeleteArticle)
			v1.GET("/tags",controllers.GetTags)
		}
	}

	return router
}
