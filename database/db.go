package database

import "github.com/jackc/pgx/v4/pgxpool"

// Storage DB
type Storage struct {
	DB      *pgxpool.Pool
	DBSlave *pgxpool.Pool
}

// New DB
func New(dbMaster, dbSlave *pgxpool.Pool) *Storage {
	return &Storage{
		DB:      dbMaster,
		DBSlave: dbSlave,
	}
}
