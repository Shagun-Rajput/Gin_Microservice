package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/service"
)

// PingHandler handles the /ping endpoint.
func PingHandler(c *gin.Context) {
    response := service.PingService()
    c.JSON(http.StatusOK, response)
}