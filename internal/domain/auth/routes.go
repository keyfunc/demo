package auth

import (
	"log/slog"
	"mall/internal/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthDeps struct {
	Cfg    *config.Config
	Logger *slog.Logger
	DB     *pgxpool.Pool
	Mux    *http.ServeMux
}

func NewRoutes(deps *AuthDeps) {
	repository := NewRepository(RepositoryDeps{DB: deps.DB})
	service := NewService(ServiceDeps{Repo: repository})
	handler := NewHandler(HandlerDeps{Logger: deps.Logger, Service: service})

	// auth
	deps.Mux.HandleFunc("POST /api/v1/auth/login", handler.Login)
	// deps.Mux.HandleFunc("POST /api/v1/auth/refresh", handler.Create)
	// deps.Mux.HandleFunc("POST /api/v1/auth/logout", handler.Create)
}
