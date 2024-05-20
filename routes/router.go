package routes

import (
	"github.com/gin-gonic/gin"
	"post-service/configs"
	"post-service/handler"
	"post-service/services"
)

func Router(router *gin.Engine) {
	dsn := configs.GetEnv()

	svc, err := service.NewService(dsn)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	h := handler.NewHandler(svc)

	post := router.Group("/posts")
	{
		post.POST("/", h.NewPost)
		post.GET("/", h.GetPosts)
		post.GET("/:id", h.GetPostById)
		post.PUT("/", h.UpdatePost)
		post.DELETE("/:id", h.DeletePostById)

	}
}
