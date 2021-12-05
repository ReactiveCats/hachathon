import { createTheme } from '@mui/material/styles';

export const theme = createTheme({ palette: {
    gray: { main: '#808080' },
    gradientBlue : { main: 'linear-gradient(81.29deg, #4A7DFF 12.65%, #3F51B5 75.9%)' },
    white: { main: '#ffffff' },
	navbarBackground: { main: 'rgb(75, 125, 255)' },
    taskBgLowImportant: { main: 'rgba(0, 255, 0, 0.2)' },
    taskBgMediumImportant: { main: 'rgba(255, 255, 0, 0.2)' },
    taskBgHighImportant: { main: 'rgba(255, 0, 0, 0.2)' },
    accent: { red: '#FF6782', green: '#9DE092', yellow: '#FFFF8D' }
} });
