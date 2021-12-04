import { createContext, useContext, useReducer } from 'react';
import { taskListService } from '../services/task-list-service';
import { makeDispatchWithEffects } from '../../shared/make-dispatch-with-effects';

const initialState = {
  items: [],
  loading: false,
  error: null,
};

const TaskListContext = createContext();

export function useTaskListContext() {
  return useContext(TaskListContext);
}

export const TASK_LIST_LOAD_ITEMS = 'task list load items';
export const TASK_LIST_SET_ITEMS = 'task list set items';
export const TASK_LIST_ADD_ITEM = 'task list add item';
const TASK_LIST_LOAD_ITEMS_SUCCESS = 'task list load items success';
const TASK_LIST_LOAD_ITEMS_FAILURE = 'task list load items failure';

function loadItems(data, dispatch) {
  taskListService
    .loadItems()
    .then((items) => {
      dispatch({ type: TASK_LIST_LOAD_ITEMS_SUCCESS, data: items });
    })
    .catch((err) => {
      dispatch({ type: TASK_LIST_LOAD_ITEMS_FAILURE, data: err });
    });
}

const taskListEffects = {
  [TASK_LIST_LOAD_ITEMS]: loadItems,
};

function taskListReducer(state, action) {
  const { type, data = {} } = action;

  switch (type) {
    case TASK_LIST_LOAD_ITEMS:
      return {
        ...state,
        loading: true,
        error: null,
      };
    case TASK_LIST_SET_ITEMS:
    case TASK_LIST_LOAD_ITEMS_SUCCESS:
      return {
        items: data,
        loading: false,
        error: null,
      };
    case TASK_LIST_LOAD_ITEMS_FAILURE:
      return {
        ...state,
        loading: false,
        error: data,
      };
    case TASK_LIST_ADD_ITEM:
      return {
        ...state,
        items: [...state.items, data],
      };
    default:
      throw new Error('Unsupported action type');
  }
}

export function TaskListProvider({ children }) {
  const [state, dispatch] = useReducer(taskListReducer, initialState);

  const dispatchWithEffects = makeDispatchWithEffects(
    dispatch,
    taskListEffects,
  );

  return (
    <TaskListContext.Provider value={[state, dispatchWithEffects]}>
      {children}
    </TaskListContext.Provider>
  );
}
