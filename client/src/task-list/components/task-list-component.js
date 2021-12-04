import { useEffect } from 'react';
import {
  Button,
  Container,
  List,
} from '@mui/material';
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
  background: theme => theme.palette.gradientBlue.main,
  height: '54px'
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
        <Box mt={5} sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
          <List
            sx={{
              display: 'flex',
              flexDirection: 'column',
              gap: 3,
              padding: 0,
            }}
          >
            {state.items.map(
              ({ title, icon, description, importance }, index) => {
                return (
                  <TaskListItem
                    index={index}
                    title={title}
                    icon={icon}
                    description={description}
                    importance={importance}
                  />
                );
              },
            )}
          </List>
          <Button variant="contained" onClick={handleAddTask} fullWidth sx={addButtonStyle}>
            Add task
          </Button>
        </Box>
      </Container>
      <TaskModal onSave={handleSave}></TaskModal>
    </>
  );
}
