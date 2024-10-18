import { useEffect, useState } from "react";
import Loader from './components/Loader';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';
import { fetchTodos, addTodo, deleteTodo, completeTodo } from './services/api';

import './App.css'
export type Todo = {
    Id: string;
    Title: string;
    Completed: boolean;
    CreatedAt: Date;
    UpdatedAt: Date;
};

export default function App() {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const loadTodos = async () => {
            setLoading(true);
            try {
                const data = await fetchTodos();
                setTodos(Array.isArray(data.todos) ? data.todos : []);
            } catch (error) {
                console.error("Error fetching todos:", error);
            } finally {
                setLoading(false);
            }
        };
        loadTodos();
    }, []);

    const handleAddTodo = async (title: string) => {
        const newTodo = await addTodo(title);
        setTodos(prevTodos => [...prevTodos, newTodo]);
    };

    const handleDeleteTodo = async (id: string) => {
        await deleteTodo(id);
        setTodos(prevTodos => prevTodos.filter(todo => todo.Id !== id));
    };

    const handleCompleteTodo = async (id: string) => {
        const todoToUpdate = todos.find(todo => todo.Id === id);
        if (todoToUpdate) {
            await completeTodo(id, todoToUpdate.Completed);
            setTodos(prevTodos =>
                prevTodos.map(todo =>
                    todo.Id === id ? { ...todo, Completed: !todo.Completed } : todo
                )
            );
        }
    };

    return (
        <div className="App">
            <h1>Todos</h1>
            <TodoForm onAdd={handleAddTodo} />
            {loading ? (
                <Loader />
            ) : (
                <TodoList todos={todos} onDelete={handleDeleteTodo} onComplete={handleCompleteTodo} />
            )}
        </div>
    );
}
