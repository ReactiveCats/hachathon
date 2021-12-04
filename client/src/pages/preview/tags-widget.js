import { Container } from '@mui/material';

import TagsWidget from '../../tags-widget/tags-widget';

function TaskWidgetPreviewPage() {
  return (
    <>
      <Container sx={{ margin: 2}}>
        <TagsWidget />
      </Container>
    </>
  );
}


export default TaskWidgetPreviewPage;