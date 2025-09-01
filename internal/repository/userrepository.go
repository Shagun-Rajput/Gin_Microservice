package repository


import (
    "context"
    "database/sql"
	"Gin_Microservice/internal/model"
)

// GetAllUsers fetches all users from the 'users' table.
func GetAllUsers(db *sql.DB) ([]model.User, error) {
    rows, err := db.QueryContext(context.Background(), "SELECT id, username, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var u model.User
        if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, rows.Err()
}