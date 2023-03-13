package services

import (
	"errors"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
)

type GameService struct {
	repos *repositories.Repository
}

func NewGameService(repos *repositories.Repository) *GameService {
	return &GameService{repos: repos}
}

func (s *GameService) Create(userID int, gameInput models.CreateGame) (int, error) {
	game := models.Game{
		Title:    gameInput.Title,
		Platform: gameInput.Platform,
		Status:   gameInput.Status,
		Notes:    gameInput.Notes,
	}

	return s.repos.GameInterface.Create(userID, game)
}

func (s *GameService) GetAll(userID int) ([]models.Game, error) {
	return s.repos.GameInterface.GetAll(userID)
}

func (s *GameService) GetByID(userID, gameID int) (models.Game, error) {
	return s.repos.GameInterface.GetByID(userID, gameID)
}

func (s *GameService) Update(userID int, gameInput models.UpdateGame) error {
	game, err := s.repos.GameInterface.GetByID(userID, gameInput.ID)
	if err != nil {
		return errors.New("game doesn't exist or you dont have enough rules to update it")
	}

	gameInput.ToGame(&game)

	return s.repos.GameInterface.Update(game)
}

func (s *GameService) Delete(userID, gameID int) error {
	_, err := s.repos.GameInterface.GetByID(userID, gameID)
	if err != nil {
		return errors.New("game doesn't exist or you dont have enough rules to delete it")
	}

	return s.repos.GameInterface.Delete(userID, gameID)
}
