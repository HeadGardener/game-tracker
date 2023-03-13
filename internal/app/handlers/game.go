package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) createGame(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		h.newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	var game models.CreateGame

	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := game.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	gameID, err := h.service.GameInterface.Create(userID, game)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": gameID,
	})
}

func (h *Handler) getAllGames(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		h.newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	games, err := h.service.GameInterface.GetAll(userID)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, map[string]interface{}{
		"games": games,
	})
}

func (h *Handler) getGameByID(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		h.newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	gameID, err := strconv.Atoi(chi.URLParam(r, "game_id"))
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid game_id param")
		return
	}

	game, err := h.service.GameInterface.GetByID(userID, gameID)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusOK, game)
}

func (h *Handler) updateGame(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		h.newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	gameID, err := strconv.Atoi(chi.URLParam(r, "game_id"))
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid game_id param")
		return
	}

	var gameInput models.UpdateGame

	if err := json.NewDecoder(r.Body).Decode(&gameInput); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid data to bind gameInput")
		return
	}
	gameInput.ID = gameID

	defer r.Body.Close()

	err = h.service.GameInterface.Update(userID, gameInput)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"info": fmt.Sprintf("game with id %d updated", gameID),
	})
}

func (h *Handler) deleteGame(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		h.newErrResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	gameID, err := strconv.Atoi(chi.URLParam(r, "game_id"))
	if err != nil {
		h.newErrResponse(w, http.StatusBadRequest, "invalid game_id param")
		return
	}

	if err := h.service.GameInterface.Delete(userID, gameID); err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"info": fmt.Sprintf("game with id %d deleted", gameID),
	})
}
