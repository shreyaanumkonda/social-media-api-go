package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}
type Post struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserID int `json:"user_id"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	 query:= `
	INSERT INTO posts (title, content, user_id, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id, title, content, user_id, tags, created_at, updated_at
	`
	
	err := s.db.QueryRowContext(ctx, query, post.Title, post.Content, post.UserID, pq.Array(post.Tags)).Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	fmt.Println("Creating post")
	return nil
}
