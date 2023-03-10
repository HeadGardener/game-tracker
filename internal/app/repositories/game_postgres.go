package repositories

import (
	"context"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/georgysavva/scany/v2/pgxscan"
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

func (r *GameRepository) GetAll(userID int) ([]models.Game, error) {
	var games []models.Game
	query := fmt.Sprintf(`SELECT g.id, g.title, g.platform, g.status, g.notes FROM %s
								g INNER JOIN %s ug on g.id = ug.game_id WHERE ug.user_id=$1`,
		gamesTable, usersGamesTable)

	if err := pgxscan.Select(context.Background(), r.conn, &games, query, userID); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) GetByID(userID, gameID int) (models.Game, error) {
	var game models.Game
	query := fmt.Sprintf(`SELECT g.id, g.title, g.platform, g.status, g.notes FROM %s
								g INNER JOIN %s ug on g.id = ug.game_id WHERE ug.user_id=$1 AND ug.game_id=$2`,
		gamesTable, usersGamesTable)
	if err := pgxscan.Get(context.Background(), r.conn, &game, query, userID, gameID); err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (r *GameRepository) Update(game models.Game) error {
	query := fmt.Sprintf(`UPDATE %s g SET title=$1, platform=$2, status=$3, notes=$4 WHERE g.id=$5`,
		gamesTable)

	_, err := r.conn.Exec(context.Background(), query, game.Title, game.Platform, game.Status, game.Notes, game.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) Delete(userID, gameID int) error {
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	deleteGameQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", gamesTable)
	_, err = tx.Exec(context.Background(), deleteGameQuery, gameID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	deleteUsersGameQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND game_id=$2", usersGamesTable)
	_, err = tx.Exec(context.Background(), deleteUsersGameQuery, userID, gameID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return tx.Commit(context.Background())
}
