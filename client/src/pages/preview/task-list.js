import { Container, Typography } from '@mui/material';
import { TaskList } from '../../task-list/components/task-list-component';

function TaskListPreviewPage() {
  return (
    <>
      <Container sx={{ paddingY: 4 }}>
        <Typography variant="h4" component="h1">
          Your tasks
        </Typography>
        <TaskList></TaskList>
      </Container>
    </>
  );
}

export default TaskListPreviewPage;
