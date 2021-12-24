import React from 'react';

function Todo({ todos }) {
  return todos.map((todo, index) => (
    <div
      className={todo.isComplete ? 'todo-row complete' : 'todo-row'}
      key={index}
    >
      <div key={todo.id}>{todo.name}</div>
    </div>
  ));
}

export default Todo;
