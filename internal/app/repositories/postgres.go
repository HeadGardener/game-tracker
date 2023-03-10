package repositories

import (
	"context"
	"fmt"
	"github.com/HeadGardener/game-tracker/configs"
	"github.com/jackc/pgx/v5"
)

const (
	usersTable      = "users"
	gamesTable      = "games"
	usersGamesTable = "users_games"
)

func NewDBConn(config configs.DBConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(),
		fmt.Sprintf("host=%s dbname=%s sslmode=%s", config.Host, config.DBName, config.SSLMode))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
