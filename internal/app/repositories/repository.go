package repositories

import "github.com/jackc/pgx/v5"

type Repository struct {
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{}
}
