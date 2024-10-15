package todoModels

import (
	"fmt"
	"github.com/k5sha/golang-todo-example/internal/storage/postgresql"
	"time"
)

type Todo struct {
	Id        string
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoStorage struct {
	*postgresql.Storage
}

func (s *TodoStorage) AllTodos(limit int) ([]Todo, error) {
	rows, err := s.DB.Query("SELECT id, title, completed, created_at, updated_at FROM todos LIMIT $1", limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoStorage) GetOneTodo(id string) (Todo, error) {
	var todo Todo

	err := s.DB.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id=$1", id).Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (s *TodoStorage) SaveTodo(title string) (int64, error) {
	const op = "storage.models.Todo.Save"

	stmt, err := s.DB.Prepare(`INSERT INTO todos (title, completed) VALUES ($1, FALSE) RETURNING id;`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id int64
	err = stmt.QueryRow(title).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *TodoStorage) RemoveTodo(id string) error {
	const op = "storage.models.Todo.Remove"

	stmt, err := s.DB.Prepare(`DELETE FROM todos WHERE id=$1`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TodoStorage) UpdateTodoStatus(id string, completed bool) (Todo, error) {
	const op = "storage.models.Todo.UpdateTodoStatus"

	stmt, err := s.DB.Prepare(`UPDATE todos SET completed=$1 WHERE id=$2`)
	if err != nil {
		return Todo{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(completed, id)
	if err != nil {
		return Todo{}, fmt.Errorf("%s: %w", op, err)
	}

	var todo Todo
	err = s.DB.QueryRow("SELECT id, title, completed, created_at, updated_at FROM todos WHERE id=$1", id).Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return Todo{}, fmt.Errorf("%s: %w", op, err)
	}

	return todo, nil
}
