import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import {
  HIDE_TASK_CARD,
  useTaskModalContext,
} from '../context/task-modal-context';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  p: 4,
};

export function TaskModal() {
  const { state, dispatch } = useTaskModalContext();

  const handleClose = () => {
    dispatch({ type: HIDE_TASK_CARD });
  };

  return (
    <>
      <Modal
        open={state.open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            Text in a modal
          </Typography>
        </Box>
      </Modal>
    </>
  );
}
