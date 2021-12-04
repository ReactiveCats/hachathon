import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { Typography, Button } from '@mui/material';
import { Box } from '@mui/system';
import {
  AUTH_SIGNUP,
  AUTH_LOGIN,
  useAuthContext,
} from '../context/auth-context';
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

const buttonGroupStyle = {
  display: 'flex',
  gap: 2,
};

export function Auth() {
  const [mode, setMode] = useState('login');

  const router = useRouter();

  const { state, dispatch } = useAuthContext();

  useEffect(() => {
    if (state.authorized) {
      router.back();
    }
  }, [state]);

  const handleSubmit = (data) => {
    if (mode === 'login') {
      dispatch({ type: AUTH_LOGIN, data });
    }

    if (mode === 'signup') {
      dispatch({ type: AUTH_SIGNUP, data });
    }
  };

  return (
    <>
      <Box sx={style}>
        {mode === 'login' ? (
          <>
            <Typography variant="h6" component="span">
              Log in
            </Typography>
            <AuthForm onSubmit={handleSubmit}>
              <Box sx={buttonGroupStyle}>
                <Button type="button" onClick={() => setMode('signup')}>
                  Create account
                </Button>
                <Button type="submit" variant="contained">
                  Login
                </Button>
              </Box>
            </AuthForm>
          </>
        ) : (
          <>
            <Typography variant="h6" component="span">
              Create account
            </Typography>
            <AuthForm onSubmit={handleSubmit}>
              <Box sx={buttonGroupStyle}>
                <Button type="button" onClick={() => setMode('login')}>
                  Login
                </Button>
                <Button type="submit" variant="contained">
                  Sign up
                </Button>
              </Box>
            </AuthForm>
          </>
        )}
      </Box>
    </>
  );
}
