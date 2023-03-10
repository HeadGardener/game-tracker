package repositories

import (
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
)

type Authorization interface {
	Create(user models.User) (int, error)
	GetUser(user *models.User) error
}

type GameInterface interface {
	Create(userID int, game models.Game) (int, error)
	GetAll(userID int) ([]models.Game, error)
	GetByID(userID, gameID int) (models.Game, error)
	Update(game models.Game) error
	Delete(userID, gameID int) error
}

type Repository struct {
	Authorization
	GameInterface
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(conn),
		GameInterface: NewGameRepository(conn),
	}
}
