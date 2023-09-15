package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	TakeData(times time.Time, auth string, headers, body []byte) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (*repository, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) TakeData(times time.Time, auth string, headers, body []byte) error {
	tableDataQuery := "INSERT INTO table_data (headers, body) VALUES ($1, $2)"
	_, err := r.db.Exec(tableDataQuery, headers, body)
	if err != nil {
		return fmt.Errorf("failed to insert data into table_data: %w", err)
	}

	var id int
	err = r.db.QueryRow("SELECT id FROM table_data ORDER BY id DESC LIMIT 1").Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to get inserted id from table_data: %w", err)
	}

	userDataQuery := "INSERT INTO user_data (time, user_id, data) VALUES ($1, $2, $3)"
	_, err = r.db.Exec(userDataQuery, times, auth, id)
	if err != nil {
		return fmt.Errorf("failed to insert data into user_data: %w", err)
	}

	return nil
}
