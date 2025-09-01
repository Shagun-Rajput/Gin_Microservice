package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/router"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8181"
    }

    r := gin.Default()
    router.SetupRoutes(r)

    log.Printf("*****Starting server on port %s...", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}