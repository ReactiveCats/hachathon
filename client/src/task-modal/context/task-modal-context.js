import { createContext, useReducer, useContext } from 'react';
import { taskModalService } from '../services/task-modal-service';
import { makeDispatchWithEffects } from '../../shared/make-dispatch-with-effects';

const initialState = {
  open: false,
  data: null,
};

const TaskModalContext = createContext();
TaskModalContext.displayName = 'TaskModalContext';

export const TASK_MODAL_OPEN = 'task modal open';
export const TASK_MODAL_HIDE = 'task modal hide';
export const TASK_MODAL_SAVE = 'task modal save';

export function useTaskModalContext() {
  return useContext(TaskModalContext);
}

function saveTask(task, dispatch) {
  if (task.id) {
    taskModalService.update(task.id, task);

    return;
  }

  taskModalService.create(task);
}

const taskModalEffects = {
  [TASK_MODAL_SAVE]: saveTask,
};

function taskModalReducer(state = initialState, action) {
  const { type, data = {} } = action;

  switch (type) {
    case TASK_MODAL_OPEN:
      return {
        ...state,
        data: {
          ...action.data,
        },
        open: true,
      };
    case TASK_MODAL_HIDE:
      return {
        ...state,
        data: null,
        open: false,
      };
    case TASK_MODAL_SAVE:
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
    <TaskModalContext.Provider value={[state, dispatchWithEffects]}>
      {children}
    </TaskModalContext.Provider>
  );
}
