package routers

import (
	"net/http"
	"rice-wine-shop/api/controllers"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	hanlderFile *controllers.ControllerSaveFile,
	user *controllers.ControllerUser,

) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.Default())
	engine.Use(gin.Recovery())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	userGroup := r.Group("/user")
	{
		userGroup.POST("/add", user.CreateUser)
	}
	fileGroup := r.Group("/files")
	{
		fileGroup.StaticFS("/export", http.Dir("publics"))
		fileGroup.POST("/upload", hanlderFile.SaveFile)
	}
	return &ApiRouter{
		Engine: engine,
	}
}
