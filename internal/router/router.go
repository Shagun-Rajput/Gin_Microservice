package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all API routes for the application.
func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        log.Println("Ping endpoint hit")
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
	})
}