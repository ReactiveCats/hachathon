import Head from 'next/head';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider } from '@emotion/react';
import { theme } from '../shared/theme';
import { createEmotionCache } from '../shared/create-emotion-cache';

function HackathonApp({ Component, pageProps }) {
  return (
    <CacheProvider value={createEmotionCache()}>
      <Head>
        <title>Hackathon</title>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </CacheProvider>
  );
}

export default HackathonApp;
