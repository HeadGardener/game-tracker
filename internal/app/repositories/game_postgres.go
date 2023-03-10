package repositories

import (
	"context"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type GameRepository struct {
	conn *pgx.Conn
}

func NewGameRepository(conn *pgx.Conn) *GameRepository {
	return &GameRepository{conn: conn}
}

func (r *GameRepository) Create(userID int, game models.Game) (int, error) {
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return 0, err
	}

	var gameID int

	createGameQuery := fmt.Sprintf(`INSERT INTO %s (title, platform, status, notes) 
											VALUES ($1, $2, $3, $4) RETURNING id`,
		gamesTable)

	row := tx.QueryRow(context.Background(), createGameQuery, game.Title, game.Platform, game.Status, game.Notes)
	if err := row.Scan(&gameID); err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	createUsersGameQuery := fmt.Sprintf("INSERT INTO %s (user_id, game_id) VALUES ($1, $2)", usersGamesTable)
	_, err = tx.Exec(context.Background(), createUsersGameQuery, userID, gameID)
	if err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	return gameID, tx.Commit(context.Background())
}
