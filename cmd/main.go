package main

import (
	"net/http"
	"os"
	"post-service/routes"
	"time"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "post-service/docs"
)

// @title Post Service API
// @version 1.0
// @description API documentation for the Post service
// @host localhost:8080
// @BasePath /api/v1

func main() {
    log.Println("Post service is running...")

	docs.SwaggerInfo.BasePath = "/api/v1"
    
    gin.SetMode(gin.ReleaseMode)

    ginEngine := gin.Default()

    ginEngine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodPatch},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           15 * time.Minute,
    }))

    routes.Router(ginEngine)
    baseURL := os.Getenv("BASE_URL")
    ginEngine.Run(baseURL)
}
