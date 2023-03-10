package models

import "errors"

type Game struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Platform string `json:"platform" db:"platform"`
	Status   string `json:"status" db:"status"`
	Notes    string `json:"notes" db:"notes"`
}

type CreateGame struct {
	Title    string `json:"title"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

type UpdateGame struct {
	ID       int     `json:"id"`
	Title    *string `json:"title"`
	Platform *string `json:"platform"`
	Status   *string `json:"status"`
	Notes    *string `json:"notes"`
}

func (g *CreateGame) Validate() error {
	if g.Title == "" {
		return errors.New("title cant be empty")
	}

	return nil
}

func (ug *UpdateGame) ToGame(game *Game) Game {
	if ug.Title != nil && game.Title != *ug.Title {
		game.Title = *ug.Title
	}

	if ug.Platform != nil && game.Platform != *ug.Platform {
		game.Platform = *ug.Platform
	}

	if ug.Status != nil && game.Status != *ug.Status {
		game.Status = *ug.Status
	}

	if ug.Notes != nil && game.Notes != *ug.Notes {
		game.Notes = *ug.Notes
	}

	return *game
}
