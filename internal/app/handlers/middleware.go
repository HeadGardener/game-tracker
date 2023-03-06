package handlers

import (
	"net/http"
)

func (h *Handler) IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*gameID := chi.URLParam(r, "game_id")
		game, err := h.service....(gameID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "game", game)
		next.ServeHTTP(w, r.WithContext(ctx))*/
	})
}
