import React, { useState } from 'react';
import { Todo } from '../App';

type TodoListProps = {
    todos: Todo[];
    onDelete: (id: string) => void;
    onComplete: (id: string) => void;
};

const TodoList: React.FC<TodoListProps> = ({ todos, onDelete, onComplete }) => {
    const [removingId, setRemovingId] = useState<string | null>(null);

    const handleDelete = (id: string) => {
        setRemovingId(id);
        setTimeout(() => {
            onDelete(id);
        }, 300);
    };

    if (!todos || todos.length === 0) {
        return <h2>Пусто</h2>;
    }

    return (
        <ul className="todo-grid">
            {todos.map((todo) => (
                <li key={todo.Id} className={removingId === todo.Id ? 'removing' : ''}>
                    <div className="todo-card">
                        <h2 className="todo-title">{todo.Title}</h2>
                        <p className="todo-status">{todo.Completed ? "Завершено" : "В процессе"}</p>
                        <div className="todo-actions">
                            <button
                                className="delete-btn"
                                aria-label="Удалить задачу"
                                onClick={() => handleDelete(todo.Id)}
                            >
                                <span>&#10006;</span>
                            </button>
                            <button
                                className="complete-btn"
                                aria-label="Завершить задачу"
                                onClick={() => onComplete(todo.Id)}
                            >
                                <span>&#10004;</span>
                            </button>
                        </div>
                    </div>
                </li>
            ))}
        </ul>
    );
};

export default TodoList;
