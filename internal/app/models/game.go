package models

import "errors"

type Game struct {
	ID       int    `json:"-"`
	Title    string `json:"title"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

type CreateGameInput struct {
	Title    string `json:"title"`
	Platform string `json:"platform"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

func (g *CreateGameInput) Validate() error {
	if g.Title == "" {
		return errors.New("title cant be empty")
	}

	return nil
}
