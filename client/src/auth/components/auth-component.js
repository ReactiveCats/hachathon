import { useRouter } from 'next/router';
import { Typography } from '@mui/material';
import { Box } from '@mui/system';
import { AUTH_LOGIN, useAuthContext } from '../context/auth-context';
import { AuthForm } from './auth-form-component';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  display: 'flex',
  flexDirection: 'column',
  gap: 2,
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  borderRadius: 2,
  p: 4,
};

export function Auth() {
  const router = useRouter();

  const { state, dispatch } = useAuthContext();

  useEffect(() => {
    if (state.authorized) {
      router.back();
    }
  }, []);

  const handleSubmit = (data) => {
    dispatch({ type: AUTH_LOGIN, data });
  };

  return (
    <>
      <Box sx={style}>
        <Typography variant="h6" component="span">
          Authorization
        </Typography>
        <AuthForm onSubmit={handleSubmit}></AuthForm>
      </Box>
    </>
  );
}
