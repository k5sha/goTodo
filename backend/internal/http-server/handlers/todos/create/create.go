package create

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"log/slog"
	"net/http"
)

type Request struct {
	Title string `json:"title"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Id     int64  `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
}

type TodoSaver interface {
	SaveTodo(title string) (int64, error)
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

		id, err := todoSaver.SaveTodo(title)
		if err != nil {
			log.Error("failed to save todo", sl.Err(err))
			ResponseError(w, r, "failed to save todo")
			return

		}

		ResponseOK(w, r, id, title)
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, id int64, title string) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{
		Status: "OK",
		Id:     id,
		Title:  title,
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, Response{
		Status: "Error",
		Error:  msg,
	})
}
