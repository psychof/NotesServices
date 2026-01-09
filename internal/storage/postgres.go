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
		return nil, fmt.Errorf("Error connext database: %s", err)
	}

	defer conn.Close(ctx)

	return &Storage{db: conn}, nil
}

func (c *Storage) AddNotes(ctx context.Context, title string, text string, timeStemp *time.Time) error {

	// stmt := c.db.QueryRow(ctx, "INSERT INTO ")

	return nil

}
