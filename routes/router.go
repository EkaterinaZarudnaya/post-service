package routes

import (
	"net/http"
	"post-service/configs"
	"post-service/handlers"
	service "post-service/services"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	dsn := configs.GetEnv()

	svc, err := service.NewService(dsn)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	h := handlers.NewPostHandler(svc)

	post := router.Group("api/posts")
	{
		post.POST("/", h.NewPost)
		post.GET("/", h.GetPosts)
		post.GET("/:id", h.GetPostById)
		post.PUT("/:id", h.UpdatePostById)
		post.DELETE("/:id", h.DeletePostById)

	}

	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome!"})
	})
}
