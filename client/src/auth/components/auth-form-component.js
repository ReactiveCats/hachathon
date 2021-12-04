import { useForm } from 'react-hook-form';
import Button from '@mui/material/Button';
import { TextField } from '@mui/material';
import { Box } from '@mui/system';

const formFieldStyle = { display: 'flex', flexDirection: 'column', gap: 1 };

export function AuthForm({ onSubmit, children }) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
          <Box sx={formFieldStyle}>
            <TextField
              label="Username"
              variant="outlined"
              {...register('username', { required: true })}
            />
            {errors.username?.type === 'required' && (
              <Box sx={{ color: 'red' }} component="span">
                Username is required
              </Box>
            )}
          </Box>

          <Box sx={{ alignSelf: 'flex-end' }}>{children}</Box>
        </Box>
      </form>
    </>
  );
}
