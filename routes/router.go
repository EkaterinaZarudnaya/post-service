package routes

import (
	"net/http"

	"github.com/ekaterinazarudnaya/post-service/configs"
	"github.com/ekaterinazarudnaya/post-service/handlers"
	service "github.com/ekaterinazarudnaya/post-service/services"

	_ "github.com/ekaterinazarudnaya/post-service/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(router *gin.Engine) {
	dsn := configs.GetEnv()

	svc, err := service.NewService(dsn)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	h := handlers.NewPostHandler(svc)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
		})

		apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		post := apiGroup.Group("/posts")
		{
			post.POST("/", h.NewPost)
			post.GET("/", h.GetPosts)
			post.GET("/:id", h.GetPostById)
			post.PUT("/:id", h.UpdatePostById)
			post.DELETE("/:id", h.DeletePostById)
		}

	}

}
