import { useEffect } from 'react';
import { Button, Container, List, Typography } from '@mui/material';
import { Box } from '@mui/system';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import { TaskModal } from '../../task-modal/components/task-modal-component';
import { TaskListItem } from './task-list-item-component';
import {
  TASK_MODAL_OPEN,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';
import {
  TASK_LIST_SAVE_ITEM,
  TASK_LIST_LOAD_ITEMS,
  useTaskListContext,
} from '../context/task-list-context';

const addButtonStyle = (theme) => ({
  background: (theme) => theme.palette.gradientBlue.main,
  height: '54px'
});

const listStyle = (theme) => ({
  display: 'flex',
  flexDirection: 'column',
  gap: 3,
  padding: 3,
  maxHeight: 'calc(100vh - 550px)',
  overflowY: 'auto',
  overflowX: 'hidden',
  '&::-webkit-scrollbar': {
    width: '3px'
  },
  '&::-webkit-scrollbar-track': {
    boxShadow: 'inset 0 0 6px rgba(0, 0, 0, 0.3)',
    borderRadius: '10px'
  },
  '&::-webkit-scrollbar-thumb': {
    backgroundColor: 'darkgrey',
  },
  [theme.breakpoints.down('sm')]: {
    maxHeight: 'calc(100vh - 430px)',
  }
})

export function TaskList() {
  const [state, dispatch] = useTaskListContext();
  const [taskModalState, taskModalDispatch] = useTaskModalContext();

  useEffect(() => {
    dispatch({ type: TASK_LIST_LOAD_ITEMS });
  }, []);

  const handleAddTask = () => {
    taskModalDispatch({
      type: TASK_MODAL_OPEN,
      data: {
        title: `Task ${state.items.length + 1}`,
      },
    });
  };

  const handleSave = (data) => {
    dispatch({ type: TASK_LIST_SAVE_ITEM, data });
  };

  return (
    <>
      <Container>
        <Box mt={5} sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
          <List
            style={state.items.length > 1 ? {minHeight: '200px'} : {minHeight: '0'}}
            sx={listStyle}
          >
            {state.items.length === 0 ? (
              <Typography align="center" variant="h5">Create Your First Task</Typography>
            ) : null}
            {state.items.map(
              ({ title, icon, description, importance, estimated }, index) => {
                return (
                  <TaskListItem
                    index={index}
                    title={title}
                    icon={icon}
                    description={description}
                    importance={importance}
                    estimated={estimated}
                  />
                );
              },
            )}
          </List>
          <Box sx={{ paddingLeft: 3, paddingRight: 3 }}>
            <Button
              variant="contained"
              onClick={handleAddTask}
              fullWidth
              sx={addButtonStyle}
            >
              Add task
            </Button>
          </Box>
        </Box>
      </Container>
      <TaskModal onSave={handleSave}></TaskModal>
    </>
  );
}
