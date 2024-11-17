package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/api/controllers"
	"rice-wine-shop/api/middlewares"
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
	rice *controllers.RiceController,

) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()
	engine.Use(gin.Logger())

	engine.Use(middlewares.CORSMiddleware())

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
	adminGroup := r.Group("/admin", middleware.AuthorizationAdmin())
	{
		typeRiceGroup := adminGroup.Group("/typeRice")
		{
			typeRiceGroup.POST("/add", typeRice.AddTypeRice)
			typeRiceGroup.GET("/list", typeRice.GetTypeRice)
		}
		riceGroup := adminGroup.Group("/rice")
		{
			riceGroup.POST("/add", rice.AddRice)
		}
	}

	fileGroup := r.Group("/files")
	{
		fileGroup.GET("/export/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			filePath := "publics/" + filename

			if _, err := http.Dir("publics").Open(filename); err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
				return
			}

			c.Writer.Header().Set("Content-Type", http.DetectContentType([]byte(filename)))

			c.Writer.Header().Set("Content-Disposition", "inline")
			c.File(filePath)
		})

		// Upload file
		fileGroup.POST("/upload", handlerFile.SaveFile)
	}
	return &ApiRouter{
		Engine: engine,
	}
}
