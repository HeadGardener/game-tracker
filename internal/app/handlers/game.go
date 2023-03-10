package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"net/http"
)

func (h *Handler) createGame(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	var game models.CreateGameInput

	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		newErrResponse(w, http.StatusBadRequest, "invalid data to bind game")
		return
	}

	defer r.Body.Close()

	gameID, err := h.service.GameInterface.Create(userID, game)
	if err != nil {
		newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"info": fmt.Sprintf("game with id %d created", gameID),
	})
}

func (h *Handler) getAllGames(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getGameByID(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateGame(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteGame(w http.ResponseWriter, r *http.Request) {

}
