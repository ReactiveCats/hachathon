import { useEffect } from 'react';
import {
  Button,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Tooltip,
} from '@mui/material';
import { Box } from '@mui/system';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import { TaskModal } from '../../task-modal/components/task-modal-component';
import {
  TASK_MODAL_OPEN,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';
import {
  TASK_LIST_SAVE_ITEM,
  TASK_LIST_LOAD_ITEMS,
  useTaskListContext,
} from '../context/task-list-context';
import { mockTask } from '../../task-modal/mock';

const listItemBoxStyle = {
  bgcolor: 'lightgreen',
  borderRadius: 2,

  '&:hover': {
    transform: 'scale(1.02)',
  },

  transition: 'ease-in-out transform 0.2s',
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

  const handleEdit = (index) => () => {
    taskModalDispatch({
      type: TASK_MODAL_OPEN,
      data: state.items[index],
    });
  };

  const handleSave = (data) => {
    dispatch({ type: TASK_LIST_SAVE_ITEM, data });
  };

  return (
    <>
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <List sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
          {state.items.map(({ title, description }, index) => (
            <Tooltip key={index} title="Click to edit" followCursor>
              <Box
                sx={listItemBoxStyle}
                aria-role="button"
                onClick={handleEdit(index)}
              >
                <ListItem>
                  <ListItemIcon>
                    <AccessTimeIcon />
                  </ListItemIcon>
                  <ListItemText sx={listItemTextStyle} primary={title} />
                </ListItem>
                <Box sx={listItemDescriptionBoxStyle} component="p">
                  {description}
                </Box>
              </Box>
            </Tooltip>
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
