package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (first_name, last_name, username, email, password)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, first_name, last_name, username, email, created_at, updated_at
	`
	err := s.db.QueryRowContext(ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	fmt.Println("Creating user")
	return nil
}
