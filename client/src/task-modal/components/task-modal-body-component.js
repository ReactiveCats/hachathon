import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { s } from '../../shared/s';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Avatar from '@mui/material/Avatar';
import Icon from '@mui/material/Icon';
import IconButton from '@mui/material/IconButton';
import {
  InputLabel,
  FormControl,
  Select,
  MenuItem,
  TextField,
} from '@mui/material';
import { Box } from '@mui/system';
import { icons, getIconById } from '../mock';

const formFieldStyle = { display: 'flex', flexDirection: 'column', gap: 1 };

export function TaskModalBody({ task, onClose, onSave }) {
  const [selectedIconId, setSelectedIconId] = useState(task.icon);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const handleSave = (data) => {
    data['estimated'] = s(data['estimated']);

    ['importance', 'urgency', 'estimated'].forEach((key) => {
      data[key] = parseInt(data[key], 10);
    });

    data['icon'] = selectedIconId;

    onSave(data);
  };

  const handleIconChange = (id) => () => {
    setSelectedIconId(id);
  };

  const { component: SelectedIconComponent } = getIconById(selectedIconId);

  return (
    <>
      <form onSubmit={handleSubmit(handleSave)}>
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
          <Stack alignItems="center" sx={{ marginBottom: 2 }}>
            <Avatar sx={{ width: 96, height: 96 }}>
              <SelectedIconComponent sx={{ width: 48, height: 48 }} />
            </Avatar>
          </Stack>
          <Stack justifyContent="center" direction="row" spacing={1}>
            {Object.values(icons).map(({ component: Component, id }) => (
              <IconButton key={id} onClick={handleIconChange(id)}>
                <Component />
              </IconButton>
            ))}
          </Stack>
          <Box sx={formFieldStyle}>
            <TextField
              label="Title"
              variant="outlined"
              {...register('title', { required: true, value: task.title })}
            />
            {errors.title?.type === 'required' && (
              <Box sx={{ color: 'red' }} component="span">
                Title is required
              </Box>
            )}
          </Box>
          <Box sx={formFieldStyle}>
            <TextField
              label="Description"
              variant="outlined"
              {...register('description', {
                value: task.description,
              })}
            />
          </Box>
          <Box sx={formFieldStyle}>
            <FormControl fullWidth>
              <InputLabel id="importance">Importance</InputLabel>
              <Select
                labelId="importance"
                label="Importance"
                defaultValue={task.importance}
                {...register('importance', { value: task.importance })}
              >
                <MenuItem value={3}>Low</MenuItem>
                <MenuItem value={5}>Medium</MenuItem>
                <MenuItem value={7}>High</MenuItem>
              </Select>
            </FormControl>
          </Box>
          <Box sx={formFieldStyle}>
            <FormControl fullWidth>
              <InputLabel id="urgency">Urgency</InputLabel>
              <Select
                labelId="urgency"
                label="Urgency"
                defaultValue={task.urgency}
                {...register('urgency', { value: task.urgency })}
              >
                <MenuItem value={3}>Low</MenuItem>
                <MenuItem value={5}>Medium</MenuItem>
                <MenuItem value={7}>High</MenuItem>
              </Select>
            </FormControl>
          </Box>
          <Box sx={formFieldStyle}>
            <TextField
              label="Estimated date"
              variant="outlined"
              {...register('estimated', {
                value: s(task.estimated),
              })}
            />
          </Box>
          <Box sx={{ display: 'flex', justifyContent: 'flex-end', gap: 1 }}>
            <Button type="button" onClick={onClose}>
              Cancel
            </Button>
            <Button type="submit">Save</Button>
          </Box>
        </Box>
      </form>
    </>
  );
}
