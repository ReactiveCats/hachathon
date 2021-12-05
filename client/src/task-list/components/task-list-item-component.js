import { useEffect, useState } from 'react';
import {
  Box,
  ListItem,
  ListItemIcon,
  ListItemText,
  Tooltip,
  LinearProgress,
} from '@mui/material';
import { getIconById } from '../../task-modal/mock';
import {
  TASK_MODAL_OPEN,
  useTaskModalContext,
} from '../../task-modal/context/task-modal-context';
import {
  useTaskListContext,
} from '../context/task-list-context';

const listItemBoxStyle = {
  bgcolor: 'lightgreen',
  borderRadius: '8px 8px 0px 0px',
  boxShadow: 8,
  cursor: 'pointer',

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
  marginBottom: 2,
  marginTop: '-8px',
  display: '-webkit-box',
  '-webkit-line-clamp': '3',
  '-webkit-box-orient': 'vertical',
  overflow: 'hidden',
};

const progressStyles = (theme) => ({
  height: '5px',
  backgroundColor: theme.palette.accent.green,
  position: 'relative',
  '& span': {
    backgroundColor: theme.palette.accent.red
  },
  '&:after': {
    content: '""',
    position: 'absolute',
    height: '5px',
    right: 0,
    backgroundColor: theme.palette.accent.yellow,
    width: 'var(--estimatedWidth)',
    zIndex: -1
  }
})

export function TaskListItem({ index, title, icon, description, importance, estimated = 0, deadline = 'Sun, 5 Dec 2021 15:00:00' }) {
  const [state, dispatch] = useTaskListContext();
  const [taskModalState, taskModalDispatch] = useTaskModalContext();

  const [estimatedLineWidth, setEstimatedLineWidth] = useState(null)
  const [progress, setProgress] = useState(0);
  // Обновление прогресса каждую секунду
  useEffect(() => {
    const period = new Date(deadline) - new Date('Sat, 4 Dec 2021 23:40:00');
    console.log(estimated, period);
    setEstimatedLineWidth((estimated * 1000 / period) * 100)
    const interval = setInterval(() => {
      // TODO: remove hardcoded deadline

      const now = new Date(Date.now()) - new Date('Sat, 4 Dec 2021 23:40:00');

      const currentProgress = now <= period ? (now / period) * 100 : 100;
      setProgress(currentProgress);
      if (currentProgress === 100) clearInterval(interval);
    }, 1000);
    return () => clearInterval(interval);
  },[deadline, estimated]);

  const handleEdit = (index) => () => {
    taskModalDispatch({
      type: TASK_MODAL_OPEN,
      data: state.items[index],
    });
  };

  const { component: IconComponent } = getIconById(icon);

  const itemBgColor = {
    backgroundColor: (theme) =>
      importance >= 7
        ? theme.palette.taskBgHighImportant.main
        : importance >= 5
        ? theme.palette.taskBgMediumImportant.main
        : theme.palette.taskBgLowImportant.main,
  };

  setInterval(() => {}, 1000);

  return (
    <Tooltip key={index} title="Click to edit" followCursor>
      <Box
        sx={{ ...listItemBoxStyle, ...itemBgColor }}
        ariaRole="button"
        onClick={handleEdit(index)}
      >
        <ListItem>
          <ListItemIcon>
            <IconComponent />
          </ListItemIcon>
          <ListItemText sx={listItemTextStyle} primary={title} />
        </ListItem>
        <Box sx={listItemDescriptionBoxStyle} component="p">
          {description}
        </Box>
        {deadline &&
          <LinearProgress variant="determinate" value={progress} style={{'--estimatedWidth': `${estimatedLineWidth}%`}} sx={progressStyles} />
        }
      </Box>
    </Tooltip>
  );
}
