package handlers

import (
	"encoding/json"
	"github.com/HeadGardener/game-tracker/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handler struct {
	service   *services.Service
	errLogger *zap.Logger
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service:   service,
		errLogger: newLogger(),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// auth
	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	// game
	r.Route("/game-tracker", func(r chi.Router) {
		r.Use(h.IdentifyUser)
		r.Post("/", h.createGame)
		r.Get("/", h.getAllGames)
		r.Get("/{game_id}", h.getGameByID)
		r.Put("/{game_id}", h.updateGame)
		r.Delete("/{game_id}", h.deleteGame)
	})

	return r
}

func newLogger() *zap.Logger {
	rawJSON := []byte(`{
	  "level": "error",
	  "encoding": "json",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()
	return logger
}
