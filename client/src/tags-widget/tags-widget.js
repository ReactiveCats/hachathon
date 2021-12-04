import React, { useEffect, useState } from 'react';

import { CircularProgress, Box, Grid, Typography, TextField, Chip, Paper, IconButton, autocompleteClasses } from '@mui/material';

import { wrappedFetch } from '../shared/wrapped-fetch';

const API = '/api/tags';

function TagsWidget() {
  const [tags, setTags] = useState([]);
  const [title, setTitle] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true);

    fetchTags();
  }, []);

  return (
    <Paper sx={{
      p: 2,
      width: 240,
      height: 280,
    }}>
      {
        isLoading ? <CircularProgress /> : (
          <Grid container spacing={1}>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Add Tag"
                size="small"
                variant="standard"
                value={title}
                onChange={handleChange}
                onBlur={handleCreate}
                onKeyDown={handleKeyDown}
              />
            </Grid>
            <Grid item xs={12}>
              <Box
                sx={{
                  display: 'flex',
                  flexWrap: 'wrap',
                  maxHeight: '200px',
                  overflow: 'auto',
                }}
              >
                {tags.map(tag => (
                  <Chip
                    sx={{ marginRight: 1, marginBottom: 1 }}
                    size="small"
                    key={tag.id}
                    label={tag.title}
                    onDelete={() => handleDelete(tag)}
                  />
                ))}
              </Box>
            </Grid>
          </Grid>
        )
      }

    </Paper>
  );

  async function fetchTags() {
    const response = await wrappedFetch(API);
    setTags(response);
    setIsLoading(false);
  }

  function handleChange(e) {
    const { value } = e.target;
    setTitle(value);
  }

  function handleKeyDown(e) {
    if (e.key !== 'Enter') {
      return;
    }

    handleCreate();
  }

  function handleCreate() {
    if (!title) {
      return;
    }

    wrappedFetch(API, {
      method: 'PUT',
      body: {
        title: value
      },
    });

    setTitle('');
  }

  function handleDelete(tag) {
    wrappedFetch(API, {
      method: 'DELETE',
      body: tag,
    });
  }
}

export default TagsWidget;