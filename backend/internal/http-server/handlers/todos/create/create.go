package create

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"github.com/k5sha/golang-todo-example/internal/storage/models/todoModels"
	"log/slog"
	"net/http"
)

type Request struct {
	Title string `json:"title"`
}

type Response struct {
	Status string           `json:"status"`
	Error  string           `json:"error,omitempty"`
	Todo   *todoModels.Todo `json:"todo,omitempty"`
}

type TodoSaver interface {
	SaveTodo(title string) (todoModels.Todo, error)
}

func New(log *slog.Logger, todoSaver TodoSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.todos.create.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			ResponseError(w, r, "failed to decode request")

			return
		}

		log.Info("request body decoded", slog.Any("request", req))
		if req.Title == "" {
			ResponseError(w, r, "field Title is a required field")

			return
		}

		title := req.Title

		todo, err := todoSaver.SaveTodo(title)
		if err != nil {
			log.Error("failed to save todo", sl.Err(err))
			ResponseError(w, r, "failed to save todo")
			return

		}

		ResponseOK(w, r, todo)
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, todo todoModels.Todo) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{
		Status: "OK",
		Todo:   &todo,
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, Response{
		Status: "Error",
		Error:  msg,
	})
}
