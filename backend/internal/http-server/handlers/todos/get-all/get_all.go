package getAll

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"github.com/k5sha/golang-todo-example/internal/storage/models/todoModels"
	"log/slog"
	"net/http"
	"strconv"
)

type Response struct {
	Status string            `json:"status"`
	Error  string            `json:"error,omitempty"`
	Todos  []todoModels.Todo `json:"todos"`
}

type TodosGetter interface {
	AllTodos(limit int) ([]todoModels.Todo, error)
}

func New(log *slog.Logger, todosGetter TodosGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.todos.getAll.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		limit, err := getLimit(r)
		if err != nil {
			ResponseError(w, r, err.Error())
			return
		}

		todos, err := todosGetter.AllTodos(limit)
		if err != nil {
			log.Error("failed to get todos", sl.Err(err))
			ResponseError(w, r, "failed to get todos")
			return

		}

		ResponseOK(w, r, todos)
	}
}

func getLimit(r *http.Request) (int, error) {
	limit := r.URL.Query().Get("limit")

	if limit == "" {
		return 50, nil
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		return 0, fmt.Errorf("check your limit param")
	}

	if limitNum <= 0 || limitNum > 128 {
		return 0, fmt.Errorf("limit must be between 1 and 128")
	}

	return limitNum, nil
}

func ResponseOK(w http.ResponseWriter, r *http.Request, todos []todoModels.Todo) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{
		Status: "OK",
		Todos:  todos,
	})
}

func ResponseError(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, Response{
		Status: "Error",
		Error:  msg,
	})
}
