import { createContext, useReducer, useContext } from 'react';

const initialState = {
  open: false,
  data: null,
};

const TaskModalContext = createContext();
TaskModalContext.displayName = 'TaskModalContext';

export const OPEN_TASK_CARD = 'open';
export const HIDE_TASK_CARD = 'hide';

export const useTaskModalContext = () => useContext(TaskModalContext);

function taskModalReducer(state = initialState, action) {
  switch (action.type) {
    case OPEN_TASK_CARD:
      return {
        ...state,
        data: {
          ...action.data,
        },
        open: true,
      };
    case HIDE_TASK_CARD:
      return {
        ...state,
        data: null,
        open: false,
      };
    default:
      throw new Error('Unsupported action type');
  }
}

export function TaskModalProvider({ children }) {
  const [state, dispatch] = useReducer(taskModalReducer, initialState);

  return (
    <TaskModalContext.Provider value={{ state, dispatch }}>
      {children}
    </TaskModalContext.Provider>
  );
}
