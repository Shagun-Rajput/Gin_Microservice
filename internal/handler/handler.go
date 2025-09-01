package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/service"
	"Gin_Microservice/internal/config"
)

var AppConfig *config.Config

// SetConfig sets the application config for handlers.
func SetConfig(cfg *config.Config) {
    AppConfig = cfg
}

// GetAllUsersHandler handles GET /users requests.
func GetAllUsersHandler(c *gin.Context) {
    users, err := service.GetAllUsersService(AppConfig.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}