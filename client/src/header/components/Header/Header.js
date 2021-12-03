import {
  AppBar,
  Toolbar,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from '@mui/material';

function Header() {
  return (
    <header>
      {/* Dropdown */}
      <AppBar position="static">
        <Toolbar>
          <FormControl>
            <InputLabel id="view-label">View</InputLabel>
            <Select
              labelId="view-label"
              id="view-select"
              value={0}
              label="View"
              // onChange={handleChange}
            >
              <MenuItem value={0}>Table</MenuItem>
              <MenuItem value={1}>List</MenuItem>
            </Select>
          </FormControl>
        </Toolbar>
      </AppBar>

      {/* Filters */}

      {/* Add */}
      {/* dispatch(action: open) */}
    </header>
  );
}

export default Header;
