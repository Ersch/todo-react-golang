import classes from "./AddTodo.module.css";
import { useRef, useState } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'

function AddTodo() {
    const [totoIsValid, setTodoIsValid] = useState(true);
    const titleRef = useRef();

    const submitHandler = (event) => {
        event.preventDefault();
        const enteredTitle = titleRef.current.value;
        console.log(enteredTitle);
    }

    return <div className={classes.add__todo}>
        <form className={classes.add__todo__container}>
            <input ref={titleRef}
                className={classes.add__todo__input}
                type="text"
                id="title"
                name="title"
                placeholder="Add todo"
                onSubmit={submitHandler} />
            <button className={classes.add__todo__button} type="submit">
            <FontAwesomeIcon className={classes.todo__action}  icon={icon({ name: 'check' })} size='sm' />
            </button>
            {!totoIsValid && <p>Todo should be at least 5 characters</p>}
        </form>
    </div>


}

export default AddTodo;