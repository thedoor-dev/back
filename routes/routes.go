package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thedoor-dev/back/logs"
	"github.com/thedoor-dev/back/services"
)

func Init() *gin.Engine {
	// e := gin.New()
	e := gin.Default()
	e.Use(logs.GinLogger(), logs.GinRecovery(true))

	api := e.Group("/api")
	{
		api.POST("/pagelen", services.PostLen)
		api.POST("/post", services.PostList)
		api.POST("/post/:id", services.PostOne)
		api.POST("/signin", services.Signin)

		admin := api.Group("/admin")
		{
			admin.Use(services.JWTCheck())
			admin.POST("/img", services.ImgUpload)
			admin.POST("/postnew", services.PostNew)
			admin.POST("/post/:id", services.PostOne)
		}
	}
	return e
}
