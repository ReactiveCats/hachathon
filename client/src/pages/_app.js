import Head from 'next/head';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider } from '@emotion/react';
import { theme } from '../shared/theme';
import { createEmotionCache } from '../shared/create-emotion-cache';
import { AuthProvider } from '../auth/context/auth-context';
import { AuthGuard } from '../auth/guards/auth-guard';
import { TaskModalProvider } from '../task-modal/context/task-modal-context';
import { TaskListProvider } from '../task-list/context/task-list-context';

function HackathonApp({ Component, pageProps }) {
  return (
    <CacheProvider value={createEmotionCache()}>
      <Head>
        <title>Hackathon</title>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <AuthProvider>
          <AuthGuard>
            <TaskListProvider>
              <TaskModalProvider>
                <Component {...pageProps} />
              </TaskModalProvider>
            </TaskListProvider>
          </AuthGuard>
        </AuthProvider>
      </ThemeProvider>
    </CacheProvider>
  );
}

export default HackathonApp;
