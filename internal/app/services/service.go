package services

import (
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
)

type Authorization interface {
	Create(userInput models.RegUserInput) (int, error)
	GenerateToken(userInput models.LogUserInput) (string, error)
	ParseToken(accessToken string) (int, error)
}

type GameInterface interface {
	Create(userID int, gameInput models.CreateGameInput) (int, error)
}

type Service struct {
	Authorization
	GameInterface
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		GameInterface: NewGameService(repos),
	}
}
