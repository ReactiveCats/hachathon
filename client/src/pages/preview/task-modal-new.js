import { useState } from 'react';
import Button from '@mui/material/Button';
import TaskModal from '../../task-modal-new/task-modal-new';

function TaskModalNewPreviewPage() {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <>
            <Button onClick={handleOpen}>Open modal</Button>
            <TaskModal open={isOpen} onClose={handleClose} />
        </>
    )

    function handleOpen() {
        setIsOpen(true);
    }

    function handleClose() {
        setIsOpen(false);
    }
}

export default TaskModalNewPreviewPage;