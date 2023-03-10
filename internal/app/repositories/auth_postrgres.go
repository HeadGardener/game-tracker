package repositories

import (
	"context"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type AuthRepository struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) *AuthRepository {
	return &AuthRepository{conn: conn}
}

func (r *AuthRepository) Create(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		usersTable)

	err := r.conn.QueryRow(context.Background(), query, user.Name, user.Username, user.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(user *models.User) error {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	err := r.conn.QueryRow(context.Background(), query, user.Username, user.PasswordHash).Scan(&id)
	user.ID = id

	return err
}
