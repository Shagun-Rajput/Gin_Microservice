package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// SetupRoutes sets up all API routes for the application.
func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
	})
}