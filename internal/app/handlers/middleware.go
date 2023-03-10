package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	userCtx = "userID"
)

func (h *Handler) IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			errMsg := "empty auth header"
			h.newErrResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			h.newErrResponse(w, http.StatusBadRequest, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			h.newErrResponse(w, http.StatusBadRequest, "jwt token is empty")
		}

		userID, err := h.service.Authorization.ParseToken(headerParts[1])
		if err != nil {
			h.newErrResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserID(r *http.Request) (int, error) {
	id := r.Context().Value(userCtx)

	if id == nil {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
