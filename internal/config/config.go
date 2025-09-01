package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DB         *sql.DB
}

// NewConfig creates a new Config instance from environment variables.
func NewConfig() *Config {
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
    }
}

// InitDB initializes the database connection and stores it in the Config.
func (cfg *Config) InitDB() error {
	log.Println("Initializing database connection...")
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )
    db, err := sql.Open("postgres", connStr)
    if err != nil {
		log.Fatal("Failed to open database connection:", err)
        return err
    }
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    cfg.DB = db
    return db.Ping()
}