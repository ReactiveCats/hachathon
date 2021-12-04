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
  paddingBottom: 2,
  margin: 0,
  marginTop: '-8px',
};

export function TaskListItem({ index, title, icon, description, importance }) {
  const [progress, setProgress] = useState(0);

  // Обновление прогресса каждую секунду
  useEffect(() => {
    const interval = setInterval(() => {
      // TODO: remove hardcoded deadline
      const deadline = 'Sun, 5 Dec 2021 15:00:00';

      const period = new Date(deadline) - new Date('Sat, 4 Dec 2021 23:40:00');
      const now = new Date(Date.now()) - new Date('Sat, 4 Dec 2021 23:40:00');

      const currentProgress = now <= period ? (now / period) * 100 : 100;
      setProgress(currentProgress);
      if (currentProgress === 100) clearInterval(interval);
    }, 1000);
    return () => clearInterval(interval);
  });

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
        <LinearProgress variant="determinate" value={progress} />
      </Box>
    </Tooltip>
  );
}
