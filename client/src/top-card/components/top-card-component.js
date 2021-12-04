import { Box, Container, Grid, Typography } from '@mui/material';
import tasks from '../../shared/mock-tasks.json';
import icons from '../../shared/icons';

const task = tasks[0];
const Icon = icons[task.icon];

/* Styles */

const boxStyles = {
  background: theme => theme.palette.gradientBlue.main,
  height: '375px'
}

const iconStyle = (theme) => ({
  fontSize: 70,
  [theme.breakpoints.up('sm')]: {
    fontSize: 215
  }
})

const containerStyle = {
  height: '100%',
  paddingTop: '100px'
}

const TopCardComponent = () => {
  return (
    <Box
      sx={boxStyles}
    >
      <Container
        sx={containerStyle}
      >
        <Grid container>
          <Grid item xs={3}>
            <Icon color='white' sx={iconStyle} />
          </Grid>
          <Grid item xs={9}>
            <Typography variant="h3" color={theme => theme.palette.white.main}>
              {tasks[0].title}
            </Typography>
            <Typography variant="body1" color={theme => theme.palette.white.main}>
              {tasks[0].description}
            </Typography>

          </Grid>
        </Grid>
      </Container>
    </Box>
  );
};

export default TopCardComponent;