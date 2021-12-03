import { createContext, useReducer, useContext } from 'react';
import { taskModalService } from '../services/task-modal-service';
import { makeDispatchWithEffects } from '../../shared/make-dispatch-with-effects';

const initialState = {
  open: false,
  data: null,
};

const TaskModalContext = createContext();
TaskModalContext.displayName = 'TaskModalContext';

export const OPEN_TASK_CARD = 'open';
export const HIDE_TASK_CARD = 'hide';
export const UPDATE_TASK_CARD = 'update';

export function useTaskModalContext() {
  return useContext(TaskModalContext);
}

function updateTask(task, dispatch) {
  taskModalService.update(task.id, task);
}

const taskModalEffects = {
  [UPDATE_TASK_CARD]: updateTask,
};

function taskModalReducer(state = initialState, action) {
  const { type, data = {} } = action;

  switch (type) {
    case OPEN_TASK_CARD:
      return {
        data,
        open: true,
      };
    case HIDE_TASK_CARD:
      return {
        ...state,
        open: false,
      };
    case UPDATE_TASK_CARD:
      return {
        ...state,
        data: {
          ...state.data,
          data,
        },
      };
    default:
      throw new Error('Unsupported action type');
  }
}

export function TaskModalProvider({ children }) {
  const [state, dispatch] = useReducer(taskModalReducer, initialState);

  const dispatchWithEffects = makeDispatchWithEffects(
    dispatch,
    taskModalEffects,
  );

  return (
    <TaskModalContext.Provider value={{ state, dispatch: dispatchWithEffects }}>
      {children}
    </TaskModalContext.Provider>
  );
}
