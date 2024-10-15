package httpServer

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k5sha/golang-todo-example/internal/http-server/handlers/todos/create"
	getAll "github.com/k5sha/golang-todo-example/internal/http-server/handlers/todos/get-all"
	getOne "github.com/k5sha/golang-todo-example/internal/http-server/handlers/todos/get-one"
	"github.com/k5sha/golang-todo-example/internal/http-server/handlers/todos/remove"
	"github.com/k5sha/golang-todo-example/internal/http-server/handlers/todos/update"
	"github.com/k5sha/golang-todo-example/internal/http-server/middleware/logger"
	"github.com/k5sha/golang-todo-example/internal/storage/models/todoModels"
	"github.com/k5sha/golang-todo-example/internal/storage/postgresql"
	"log/slog"
)

func SetupRoutes(log *slog.Logger, storage *postgresql.Storage) *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(logger.New(log))

	// Init Storage
	todoStorage := &todoModels.TodoStorage{Storage: storage}

	// Routes
	router.Route("/todo", func(r chi.Router) {
		r.Get("/", getAll.New(log, todoStorage))
		r.Get("/{id}", getOne.New(log, todoStorage))

		r.Post("/", create.New(log, todoStorage))
		r.Delete("/", remove.New(log, todoStorage))
		r.Patch("/{id}/status", update.New(log, todoStorage))

	})

	return router
}
