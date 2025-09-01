package service

import (
	"Gin_Microservice/internal/model"
	"Gin_Microservice/internal/repository"
	"database/sql"
	"log"
)

// GetAllUsersService fetches all users using the repository.
func GetAllUsersService(db *sql.DB) ([]model.User, error) {
	log.Println("Inside service : Fetching all users from the database")
    return repository.GetAllUsers(db)
}