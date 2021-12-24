import React, { useState, useEffect, useRef } from 'react';

function TodoForm(props) {
  const [input, setInput] = useState('');

  const inputRef = useRef(null); // Ref Using

  useEffect(() => {
    inputRef.current.focus();
  });

  const handleChange = (e) => {
    setInput(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: input }),
    };
    fetch('http://127.0.0.1:5000/api/todo', requestOptions)
      .then((res) => res.json())
      .then((res) => props.onSubmit(res));

    setInput('');
  };

  return (
    <form className='todo-form' onSubmit={handleSubmit}>
      <input
        placeholder='Add a todo'
        value={input}
        onChange={handleChange}
        name='name'
        className='todo-input'
        ref={inputRef}
      />
      <button onClick={handleSubmit} className='todo-button'>
        Add todo
      </button>
    </form>
  );
}
export default TodoForm;
