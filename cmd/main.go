package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "Gin_Microservice/internal/router"
    "Gin_Microservice/internal/config"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8181"
    }

    // Initialize config and DB connection
    cfg := config.NewConfig()
    if err := cfg.InitDB(); err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    r := gin.Default()
    router.SetupRoutes(r, cfg)

    log.Printf("*****Starting server on port %s...", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}