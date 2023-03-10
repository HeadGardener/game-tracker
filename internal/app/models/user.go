package models

import "errors"

type User struct {
	ID           int    `json:"-"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type RegUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *RegUserInput) Validate() error {
	if u.Name == "" || u.Username == "" || u.Password == "" {
		return errors.New("there can't be empty fields in user struct")
	}

	return nil
}

func (u *LogUserInput) Validate() error {
	if u.Username == "" || u.Password == "" {
		return errors.New("there can't be empty fields in user struct")
	}

	return nil
}
