package handler

import (
	"log"
	"net/http"

	"Gin_Microservice/internal/config"
	"Gin_Microservice/internal/service"

	"github.com/gin-gonic/gin"
)

var AppConfig *config.Config

// SetConfig sets the application config for handlers.
func SetConfig(cfg *config.Config) {
    AppConfig = cfg
}

// GetAllUsersHandler handles GET /users requests.
func GetAllUsersHandler(c *gin.Context) {
    log.Println("GetAllUsersHandler called.....")
    users, err := service.GetAllUsersService(AppConfig.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}