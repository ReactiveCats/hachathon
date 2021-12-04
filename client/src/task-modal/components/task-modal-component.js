import Box from '@mui/material/Box';
import Modal from '@mui/material/Modal';
import isEqual from '@tinkoff/utils/is/equal';
import {
  TASK_MODAL_HIDE,
  TASK_MODAL_SAVE,
  useTaskModalContext,
} from '../context/task-modal-context';
import { TaskModalBody } from './task-modal-body-component';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  borderRadius: 2,
  p: 4,
};

const taskFromData = (data) => ({
  title: '',
  description: '',
  ...data,
});

export function TaskModal({ onSave }) {
  const [state, dispatch] = useTaskModalContext();

  if (state.data === null) {
    return <></>;
  }

  const close = () => {
    dispatch({ type: TASK_MODAL_HIDE });
  };

  const handleSave = (data) => {
    if (!isEqual(data, state.data)) {
      dispatch({ type: TASK_MODAL_SAVE, data });

      if (typeof onSave === 'function') {
        onSave(data);
      }
    }

    close();
  };

  const handleClose = () => {
    close();
  };

  const task = taskFromData(state.data);

  return (
    <>
      <Modal
        open={state.open}
        onClose={handleClose}
        aria-labelledby={task.title}
        aria-describedby={task.description}
        disablePortal
      >
        <Box sx={style}>
          <TaskModalBody
            task={task}
            onClose={handleClose}
            onSave={handleSave}
          ></TaskModalBody>
        </Box>
      </Modal>
    </>
  );
}
