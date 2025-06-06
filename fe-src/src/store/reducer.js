import { combineReducers } from 'redux'

import todosReducer from '../features/todos/todosSlice'
import filtersReducer from '../features/filters/filtersSlice'
import reposReducer from "../features/repos/reposSlice";

const rootReducer = combineReducers({
    // Define a top-level state field named `todos`, handled by `todosReducer`
    todos: todosReducer,
    repos: reposReducer,
    filters: filtersReducer,
})

export default rootReducer