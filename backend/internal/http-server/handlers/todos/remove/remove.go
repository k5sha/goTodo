package remove

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"log/slog"
	"net/http"
)

type Request struct {
	Id string `json:"id"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type TodoRemover interface {
	RemoveTodo(id string) error
}

func New(log *slog.Logger, todoRemover TodoRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.todos.remove.New"

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
		if req.Id == "" {
			ResponseError(w, r, "field Id is a required field")

			return
		}

		id := req.Id

		err = todoRemover.RemoveTodo(id)
		if err != nil {
			log.Error("failed to remove todo", sl.Err(err))
			ResponseError(w, r, "failed to remove todo")
			return

		}

		ResponseOK(w, r)
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{
		Status: "OK",
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, Response{
		Status: "Error",
		Error:  msg,
	})
}
