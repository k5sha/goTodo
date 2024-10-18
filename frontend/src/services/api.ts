import axios from "axios";

const API_URL = 'http://localhost:8080/todo';

export const fetchTodos = async () => {
    const response = await axios.get(`${API_URL}`);
    return response.data;
};

export const addTodo = async (title: string) => {
    const response = await axios.post(API_URL, { title }, {
        headers: { "Content-Type": "application/json" },
    });
    return response.data.todo;
};

export const deleteTodo = async (id: string) => {
    await axios.delete(`${API_URL}/${id}/delete`);
};

export const completeTodo = async (id: string, completed: boolean) => {
    await axios.patch(`${API_URL}/${id}/status`, { Completed: !completed });
};
