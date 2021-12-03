import { useState } from "react";
import {
  AppBar,
  Toolbar,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button,
  Grid,
} from '@mui/material';

function Header() {
  const [view, setView] = useState(0);
  const [filter, setFilter] = useState(0);

  const handleViewChange = (event) => {
    setView(event.target.value);
  };

  const handleFilterChange = (event) => {
    setFilter(event.target.value);
  };

  return (
    <AppBar position="static">
      <Toolbar>
        <Grid
          container
          display="flex"
          justifyContent="space-between"
          alignItems="center"
          paddingY="16px"
          width="100%"
        >
          {/* Views */}
          <Grid item>
            <FormControl>
              <InputLabel id="view-label">View</InputLabel>
              <Select
                labelId="view-label"
                id="view-select"
                value={view}
                label="View"
                onChange={handleViewChange}
              >
                <MenuItem value={0}>Table</MenuItem>
                <MenuItem value={1}>List</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          {/* Filters */}
          <Grid item>
            <FormControl>
              <InputLabel id="filter-label">Filters</InputLabel>
              <Select
                labelId="filter-label"
                id="filter-select"
                value={filter}
                label="Filter"
                onChange={handleFilterChange}
              >
                <MenuItem value={0}>Filter 1</MenuItem>
                <MenuItem value={1}>Filter 2</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          {/* Add Task Button */}
          <Grid item>
            <Button variant="contained">+</Button>
          </Grid>
        </Grid>
      </Toolbar>
    </AppBar>
  );
}

export default Header;
