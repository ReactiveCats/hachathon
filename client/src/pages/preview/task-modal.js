import Button from '@mui/material/Button';
import { TaskModal } from '../../task-modal/components/task-modal-component';
import {
  TASK_MODAL_OPEN,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';
import { mockTask } from '../../task-modal/mock';

function TaskModalPreviewPage() {
  const [state, dispatch] = useTaskModalContext();

  const handleClick = () => {
    dispatch({ type: TASK_MODAL_OPEN, data: mockTask({ title: 'Preview' }) });
  };

  return (
    <>
      <Button onClick={handleClick}>Open task modal</Button>
      <TaskModal />
    </>
  );
}

export default TaskModalPreviewPage;
