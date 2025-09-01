package router
// ...existing code...
import (
    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/handler"
)

// SetupRoutes sets up all API routes for the application.
func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", handler.PingHandler)
}