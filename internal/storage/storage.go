package storage

import (
	"context"
	"database/sql"
)

type PostRepository interface {
	Create(context.Context, *Post) error
}

type UserRepository interface {
	Create(context.Context, *User) error
}

type Storage struct {
	Posts PostRepository
	Users UserRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db: db},
		Users: &UsersStore{db: db},
	}
}
