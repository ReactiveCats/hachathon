import React, { useState } from 'react';
import { Box, Grid, Button, Typography, Modal, TextField, Avatar, Stack, IconButton } from '@mui/material';

import { getIcons, getIconById } from '../constants/icons';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 800,
  bgcolor: 'background.paper',
  boxShadow: 24,
  p: 4,
};

const defaultData = {
  icon: '0',
  title: '',
  description: '',
  deadline: '',
  estimated_time: '',
}

function TaskModal({ open, onClose }) {
  const [data, setData] = useState(defaultData);

  const { icon, title, description, deadline, estimated_time } = data;
  const Icon = getIconById(icon).component;
  const icons = getIcons();

  return (
    <div>
      <Modal
        open={open}
        onClose={onClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography variant="h4"> Add Task </Typography>
          <Grid container spacing={2} sx={{ marginTop: 2, marginBottom: 2 }}>
            <Grid item xs={6}>
              <Stack alignItems="center" sx={{ marginBottom: 2 }}>
                <Avatar sx={{ width: 96, height: 96 }}>
                  <Icon sx={{ width: 48, height: 48 }} />
                </Avatar>
              </Stack>
              <Stack justifyContent="center" direction="row" spacing={1}>
                {
                  icons.map(icon => {
                    const { component: Component, id } = icon;
                    return (
                      <IconButton onClick={() => handleChangeIcon(id)}>
                        <Component key={id} />
                      </IconButton>
                    )
                  })
                }
              </Stack>
            </Grid>
            <Grid item xs={6}>
              <Box
                component="form"
                sx={{
                  '& .MuiTextField-root': { m: 1 },
                }}
                noValidate
                autoComplete="off"
              >
                <TextField
                  required
                  fullWidth
                  label="Title"
                  value={title}
                  onChange={(e) => handleChange(e, 'title')}
                />
                <TextField
                  multiline
                  rows={2}
                  fullWidth
                  label="Description"
                  value={description}
                  onChange={(e) => handleChange(e, 'description')}
                />
                <Stack direction="row" spacing={2}>
                  <LocalizationProvider dateAdapter={AdapterMoment}>
                    <DatePicker
                      label="Deadline"
                      value={deadline}
                      onChange={(e) => handleChange(e, 'deadline')}
                      renderInput={(params) => <TextField {...params} />}
                    />
                  </LocalizationProvider>
                  <TextField
                    fullWidth
                    label="Estimated Time"
                    value={estimated_time}
                    onChange={(e) => handleChange(e, 'estimated_time')}
                  />
                </Stack>
              </Box>
            </Grid>
          </Grid>
          <Stack direction="row" spacing={2} justifyContent="flex-end">
            <Button prim>Cancel</Button>
            <Button variant="contained">Add Task</Button>
          </Stack>
        </Box>
      </Modal>
    </div>
  );

  function handleChange(e, key) {
    const { value } = e.target;

    setData(prev => ({
      ...prev,
      [key]: value,
    }));
  };

  function handleChangeIcon(icon) {
    setData(prev => ({
      ...prev,
      icon,
    }))
  }
}

export default TaskModal;