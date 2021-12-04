import { useEffect } from 'react';
import {
  Button,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from '@mui/material';
import { Box } from '@mui/system';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
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

const listItemBoxStyle = {
  bgcolor: 'lightgreen',
  borderRadius: 2,
};

const listItemTextStyle = {
  // Заголовок
  span: { fontSize: 'h5.fontSize', marginLeft: '-24px' },
  // Описание
  p: {},
};

const listItemDescriptionBoxStyle = {
  fontSize: 'body1',
  paddingX: 2,
  paddingBottom: 2,
  margin: 0,
  marginTop: '-8px',
};

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
    dispatch({ type: TASK_LIST_ADD_ITEM, data });
  };

  return (
    <>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <List sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
          {state.items.map(({ title, description }, index) => (
            <Box sx={listItemBoxStyle}>
              <ListItem key={index}>
                <ListItemIcon>
                  <AccessTimeIcon />
                </ListItemIcon>
                <ListItemText sx={listItemTextStyle} primary={title} />
              </ListItem>
              <Box sx={listItemDescriptionBoxStyle} component="p">
                {description}
              </Box>
            </Box>
          ))}
        </List>
        <Button variant="outlined" onClick={handleAddTask} fullWidth>
          Add task
        </Button>
      </Box>
      <TaskModal onSave={handleSave}></TaskModal>
    </>
  );
}
