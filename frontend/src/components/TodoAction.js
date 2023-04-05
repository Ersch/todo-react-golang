import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import classes from "./TodoAction.module.css";

function TodoAction({ todo, onEdit, onRemove }) {
    return (
        <>
            <FontAwesomeIcon className={classes.todo__action} onClick={() => onEdit(todo)} icon={icon({ name: 'check' })} size='sm' />
            <FontAwesomeIcon className={classes.todo__action} onClick={() => onRemove(todo)} icon={icon({ name: 'trash' })} size='sm' />
        </>
    );
}

export default TodoAction;