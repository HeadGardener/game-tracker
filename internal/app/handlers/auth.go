package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user models.RegUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := user.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.Create(user)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var user models.LogUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := user.Validate(); err != nil {
		h.newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(user)
	if err != nil {
		h.newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(w, http.StatusCreated, map[string]interface{}{
		"token": fmt.Sprintf("%s", token),
	})
}
