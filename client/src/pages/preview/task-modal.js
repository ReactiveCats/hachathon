import Button from '@mui/material/Button';
import { TaskModal } from '../../task-modal/components/task-modal-component';
import {
  OPEN_TASK_CARD,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';

function TaskModalPreviewPage() {
  const { state, dispatch } = useTaskModalContext();

  const handleClick = () => {
    dispatch({ type: OPEN_TASK_CARD, data: {} });
  };

  return (
    <>
      <Button onClick={handleClick}>Open task modal</Button>
      <TaskModal />
    </>
  );
}

export default TaskModalPreviewPage;
