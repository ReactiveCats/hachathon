import Button from '@mui/material/Button';
import { useForm } from 'react-hook-form';
import {
  InputLabel,
  FormControl,
  Select,
  MenuItem,
  TextField,
} from '@mui/material';
import { Box } from '@mui/system';

const formFieldStyle = { display: 'flex', flexDirection: 'column', gap: 1 };

export function TaskModalBody({ task, onClose, onSave }) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  return (
    <>
      <form onSubmit={handleSubmit(onSave)}>
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
          <Box sx={formFieldStyle}>
            <TextField
              label="Title"
              variant="outlined"
              {...register('title', { required: true, value: task.title })}
            />
            {errors.title?.type === 'required' && (
              <span>Title is required</span>
            )}
          </Box>
          <Box sx={formFieldStyle}>
            <TextField
              label="Description"
              variant="outlined"
              {...register('description', {
                required: true,
                value: task.description,
              })}
            />
          </Box>
          <Box sx={formFieldStyle}>
            <FormControl fullWidth>
              <InputLabel id="priority">Priority</InputLabel>
              <Select
                labelId="priority"
                label="Priority"
                defaultValue={task.priority}
                {...register('priority', { value: task.priority })}
              >
                <MenuItem value="very_low">Very low</MenuItem>
                <MenuItem value="low">Low</MenuItem>
                <MenuItem value="medium">Medium</MenuItem>
                <MenuItem value="high">High</MenuItem>
                <MenuItem value="very_high">Very high</MenuItem>
              </Select>
            </FormControl>
          </Box>
          <Box sx={formFieldStyle}>
            <FormControl fullWidth>
              <InputLabel id="complexity">Complexity</InputLabel>
              <Select
                labelId="complexity"
                label="Complexity"
                defaultValue={task.complexity}
                {...register('complexity', { value: task.complexity })}
              >
                <MenuItem value="very_low">Very low</MenuItem>
                <MenuItem value="low">Low</MenuItem>
                <MenuItem value="medium">Medium</MenuItem>
                <MenuItem value="high">High</MenuItem>
                <MenuItem value="very_high">Very high</MenuItem>
              </Select>
            </FormControl>
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
