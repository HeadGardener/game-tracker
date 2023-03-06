package services

import "github.com/HeadGardener/game-tracker/internal/app/repositories"

type Service struct {
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{}
}
