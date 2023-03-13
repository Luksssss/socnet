package database

import "github.com/jackc/pgx/v4/pgxpool"

// Storage DB
type Storage struct {
	DB *pgxpool.Pool
}

// New DB
func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		DB: db,
	}
}
