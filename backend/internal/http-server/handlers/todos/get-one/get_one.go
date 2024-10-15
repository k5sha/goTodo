package getOne

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"github.com/k5sha/golang-todo-example/internal/storage/models/todoModels"
	"log/slog"
	"net/http"
)

type Response struct {
	Status string           `json:"status"`
	Error  string           `json:"error,omitempty"`
	Todo   *todoModels.Todo `json:"todo,omitempty"`
}

type TodosGetter interface {
	GetOneTodo(id string) (todoModels.Todo, error)
}

func New(log *slog.Logger, todosGetter TodosGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.todos.getOne.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id := chi.URLParam(r, "id")
		if id == "" {
			log.Info("id is empty")

			ResponseError(w, r, "invalid request")

			return
		}

		todo, err := todosGetter.GetOneTodo(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Info("todo not found", "id", id)

				ResponseError(w, r, "not found")

				return
			}
			log.Error("failed to get todos", sl.Err(err))
			ResponseError(w, r, "failed to get todos")
			return

		}

		ResponseOK(w, r, &todo)
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, todo *todoModels.Todo) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{
		Status: "OK",
		Todo:   todo,
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, Response{
		Status: "Error",
		Error:  msg,
		Todo:   nil,
	})
}
