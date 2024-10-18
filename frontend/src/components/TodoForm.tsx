import React, { useState } from "react";

interface TodoFormProps {
    onAdd: (title: string) => void;
}

const TodoForm: React.FC<TodoFormProps> = ({ onAdd }) => {
    const [input, setInput] = useState<string>('');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInput(event.target.value);
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (input.trim()) {
            onAdd(input);
            setInput('');
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                value={input}
                onChange={handleChange}
                placeholder="Enter your todo"
            />
            <button type="submit">Submit</button>
        </form>
    );
};

export default TodoForm;
