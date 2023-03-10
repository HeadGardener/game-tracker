package services

import (
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
)

type GameService struct {
	repos *repositories.Repository
}

func NewGameService(repos *repositories.Repository) *GameService {
	return &GameService{repos: repos}
}

func (s *GameService) Create(userID int, gameInput models.CreateGameInput) (int, error) {
	if err := gameInput.Validate(); err != nil {
		return 0, err
	}

	game := models.Game{
		Title:    gameInput.Title,
		Platform: gameInput.Platform,
		Status:   gameInput.Status,
		Notes:    gameInput.Notes,
	}

	return s.repos.GameInterface.Create(userID, game)
}
