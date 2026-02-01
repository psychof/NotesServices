package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(ctx context.Context, connString string) (*Storage, error) {
	conn, err := pgx.Connect(ctx, connString)

	if err != nil {
		return nil, fmt.Errorf("Error connect database: %s", err)
	}

	return &Storage{db: conn}, nil
}

func (c *Storage) AddNotes(ctx context.Context, title string, text string, timeStemp *time.Time) (int64, error) {

	var id int64

	stmt := c.db.QueryRow(ctx, "INSERT INTO notes (title, description, time_stamp) VALUES ($1, $2, $3) RETURNING id", title, text, timeStemp)

	err := stmt.Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error add notes:%s", err)
	}

	return id, nil

}

func (c *Storage) RemoveNotes(ctx context.Context, note_id int64, user_id int64) error {
	
	return nil
}
