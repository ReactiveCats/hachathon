import { Box, Container, Grid, Typography } from '@mui/material';
import tasks from '../../shared/mock-tasks.json';
import icons from '../../shared/icons';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import { convertTime } from '../../shared/convert-time';
import { useEffect } from 'react';

const task = tasks[1];
const Icon = icons[task.icon];

/* Styles */

const boxStyles = (theme) =>  ({
  background: theme => theme.palette.gradientBlue.main,
  minHeight: '265px',
  boxShadow: 10,
  [theme.breakpoints.up('sm')]: {
    minHeight: '350px'
  }
})

const iconStyle = (theme) => ({
  fontSize: 70,
  maxWidth: '100%',
  [theme.breakpoints.up('sm')]: {
    fontSize: 215
  }
})

const containerStyle = {
  height: '100%',
  paddingTop: '100px',
}

const titleStyle = (theme) => ({
  marginBottom: '10px',
  [theme.breakpoints.down('sm')]: {
    fontSize: 34
  }
})

const descriptionDesktopStyle = (theme) => ({
  [theme.breakpoints.down('sm')]: {
    display: 'none'
  }
})

const descriptionMobileStyle = (theme) => ({
  display: 'none',
  paddingBottom: '20px',
  [theme.breakpoints.down('sm')]: {
    display: 'block'
  }
})

const estimationStyle = (theme) => ({
  marginTop: '15px',
  display: 'flex',
  [theme.breakpoints.down('sm')]: {
    marginTop: 0,
  }
})

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
          <Grid item container xs={9} justifyContent="center" flexDirection="column">
            <Typography variant="h3" color={theme => theme.palette.white.main} sx={titleStyle}>
              {tasks[0].title}
            </Typography>
            <Typography variant="body1" color={theme => theme.palette.white.main} sx={descriptionDesktopStyle}>
              {tasks[0].description}
            </Typography>
            {tasks[0].estimated &&
              <Typography variant="body1" color={theme => theme.palette.white.main} sx={estimationStyle}>
                <AccessTimeIcon color='white' sx={{marginRight: 1}} /> {convertTime(parseInt(tasks[0].estimated))}
              </Typography>
            }
          </Grid>
          <Grid item sx={descriptionMobileStyle} mt={2}>
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