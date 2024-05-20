package main

import (
	"net/http"
	"os"
	"post-service/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.New()

	ginEngine = gin.Default()

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
