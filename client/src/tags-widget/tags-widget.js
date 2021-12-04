import React, { useEffect, useState } from 'react';

import { CircularProgress, Box, Grid, Typography, TextField, Chip, Paper, IconButton, autocompleteClasses } from '@mui/material';

import { wrappedFetch } from '../shared/wrapped-fetch';

const MOCKS = [
  {
    id: 0,
    title: 'f32',
  },
  {
    id: 1,
    title: 'bar',
  },
  {
    id: 2,
    title: 'foomore',
  },
  {
    id: 3,
    title: 'bsdfsdfa',
  },
  {
    id: 4,
    title: 'fooooo',
  },
  {
    id: 5,
    title: 'sdf',
  },
  {
    id: 6,
    title: 'dsffsdfd',
  },
  {
    id: 7,
    title: 'bar',
  },
  {
    id: 8,
    title: 'foomore',
  },
  {
    id: 9,
    title: 'bfofofa',
  },
  {
    id: 10,
    title: 'fooooo',
  },
  {
    id: 11,
    title: 'bfofoofofoffoofar',
  },
  {
    id: 12,
    title: 'dsffsdfd',
  },
  {
    id: 13,
    title: 'bar',
  },
  {
    id: 14,
    title: 'foomore',
  },
  {
    id: 15,
    title: 'bfofofa',
  },
  {
    id: 16,
    title: 'fooooo',
  },
  {
    id: 17,
    title: 'bfofoofofoffoofar',
  },
];

const API = '';

function TagsWidget() {
  const [tags, setTags] = useState(MOCKS);
  const [title, setTitle] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  // useEffect(() => {
  //   setIsLoading(true);

  //   fetchTags();
  // }, []);

  // if (isLoading) {
  //   return (<CircularProgress />);
  // }

  return (
    <Paper sx={{
      p: 2,
      width: 240,
      height: 280,
    }}>
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
    </Paper>
  );

  // async function fetchTags() {
  //   const response = await wrappedFetch(API);
  //   setTags(response);
  //   setIsLoading(false);
  // }

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

    console.log(title);

    // await wrappedFetch(API, {
    //   method: 'PUT',
    //   body: {
    //     title: value
    //   },
    // }); 

    setTitle('');
  }

  function handleDelete(tag) {
    console.log(tag);
    //   await wrappedFetch(API, {
    //     method: 'POST',
    //     body: tag,
    //   });
  }
}

export default TagsWidget;