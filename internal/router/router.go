package router
// ...existing code...
import (
    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/handler"
    "Gin_Microservice/internal/config"
)
// SetupRoutes sets up all API routes for the application.
func SetupRoutes(r *gin.Engine, cfg *config.Config) {
    handler.SetConfig(cfg)
    r.GET("/users", handler.GetAllUsersHandler)
}