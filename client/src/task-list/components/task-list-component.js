import { useEffect } from 'react';
import { Button, List } from '@mui/material';
import { TaskModal } from '../../task-modal/components/task-modal-component';
import {
  TASK_MODAL_OPEN,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';
import {
  TASK_LIST_ADD_ITEM,
  TASK_LIST_LOAD_ITEMS,
  useTaskListContext,
} from '../context/task-list-context';
import { mockTask } from '../../task-modal/mock';

export function TaskList() {
  const [state, dispatch] = useTaskListContext();
  const [taskModalState, taskModalDispatch] = useTaskModalContext();

  useEffect(() => {
    dispatch({ type: TASK_LIST_LOAD_ITEMS });
  }, []);

  const handleAddTask = () => {
    taskModalDispatch({
      type: TASK_MODAL_OPEN,
      data: {},
    });
  };

  const handleSave = (data) => {
    dispatch({ type: TASK_LIST_ADD_ITEM, data });
  };

  return (
    <>
      <List>{JSON.stringify(state.items, null, 2)}</List>
      <Button variant="outlined" onClick={handleAddTask}>
        Add task
      </Button>
      <TaskModal onSave={handleSave}></TaskModal>
    </>
  );
}
