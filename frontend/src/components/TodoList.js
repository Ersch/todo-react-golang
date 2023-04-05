import React, { Component, useEffect } from "react";
import classes from "./TodoList.module.css";
import TodoAction from "./TodoAction";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import AddTodo from "./AddTodo";

let endPoint = "http://localhost:9000/todos";

function ToDoList() {
    const dummy = [
        { id: 1, title: "Learn React", completed: false },
        { id: 2, title: "Learn GraphQL", completed: false },
        { id: 3, title: "Learn Apollo", completed: false },
    ];

    const [todos, setTodos] = React.useState([]);
    const [newTodo, setNewTodo] = React.useState("");

    useEffect(() => {
        fetch(endPoint)
            .then((res) => res.json())
            .then((data) => {
                setTodos(data);
            });
    }, []);

    const edit = (todo) => {
        console.log("Edit todo", todo);
    };

    const remove = (todo) => {
        console.log("Remove todo", todo);
    };

    return (
        <div>
            <div className={classes.title}>
                <FontAwesomeIcon icon={icon({ name: 'list-check' })} size='2xl' />
            </div>
            <AddTodo />
            {todos.length === 0 && <div className={classes.empty}>No todos</div>}
            {todos.length > 0 && <div>
                {todos.map((todo) => (
                    <div key={todo.id} className={classes['todo-list']}>
                        <div className={classes.todo}>
                            <div className={classes.todoTitle}>{todo.title}</div>
                            <div className={classes.todoActions}>
                                <TodoAction todo={todo} onEdit={edit} onRemove={remove} />
                            </div>
                        </div>
                    </div>
                ))}
            </div>}
        </div>
    );
}

export default ToDoList;

