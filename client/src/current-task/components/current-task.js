import tasks from '../../shared/mock-tasks.json';
import icons from '../../shared/icons.js';
import { Box, Grid } from '@mui/material';
import { useTheme } from '@mui/material/styles';

const boxStyle = {
  border: '1px solid #000',
  padding: "48px",
  margin: "4px"
};

const iconStyle = {
  fontSize: 150
}

const task = tasks[0];
const Icon = icons[task.icon];

function CurrentTask() {
  const theme = useTheme();
  console.log(theme);
  return (
    <Box sx={boxStyle}>
      <Grid>
        <Grid item>
          <Icon color='gray' fontSize='large'sx={iconStyle} />
        </Grid>
        <Grid item>123</Grid>
      </Grid>
    </Box>
  );
}

export default CurrentTask;
