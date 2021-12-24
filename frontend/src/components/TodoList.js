import React, { useState, useEffect } from 'react';
import TodoForm from './TodoForm';
import Todo from './Todo';

function TodoList() {
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    fetch('http://127.0.0.1:5000/api/todo')
      .then((response) => response.json())
      .then((data) => setTodos(data));
  }, []);
  const addTodo = (todo) => {
    const newTodos = [todo, ...todos];
    setTodos(newTodos);
  };

  return (
    <div>
      <h1>What's the plan for today?</h1>
      <TodoForm onSubmit={addTodo} />
      <Todo todos={todos} />
    </div>
  );
}

export default TodoList;
