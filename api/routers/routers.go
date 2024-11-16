package routers

import (
	"net/http"
	"rice-wine-shop/api/controllers"
	"rice-wine-shop/api/middlewares"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	handlerFile *controllers.ControllerSaveFile,
	user *controllers.UserController,
	auth *controllers.AuthController,
	middleware *middlewares.Middleware,
	typeRice *controllers.TypeRiceController,

) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())
	engine.Use(gin.Recovery())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/add", user.CreateUser)
	r.POST("/login", auth.Login)
	checkGroup := r.Group("/check", middleware.Authorization())
	{
		checkGroup.POST("/", middleware.CheckToken)

	}
	userGroup := r.Group("/user", middleware.Authorization())
	{
		userGroup.GET("/profile/", user.GetUser)
	}
	typeRiceGroup := r.Group("/typeRice", middleware.Authorization())
	{
		typeRiceGroup.POST("/add", typeRice.AddTypeRice)
	}
	fileGroup := r.Group("/files")
	{
		fileGroup.StaticFS("/export", http.Dir("publics"))
		fileGroup.POST("/upload", handlerFile.SaveFile)
	}
	return &ApiRouter{
		Engine: engine,
	}
}
