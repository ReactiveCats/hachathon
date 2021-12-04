import { AppBar, Toolbar, Box, Typography, Container } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

const containerStyles = {
  display: 'flex',
  paddingTop: '12px',
  paddingBottom: '8px',
};

const mobileLogoStyles = (theme) => ({
  width: '55px',
  height: '55px',
  [theme.breakpoints.up('sm')]: {
    display: 'none',
  },
});

const desktopLogoStyles = (theme) => ({
  width: '190px',
  height: '55px',
  display: 'none',
  [theme.breakpoints.up('sm')]: {
    display: 'block',
  },
});

const accountStyles = {
  marginLeft: 'auto',
  display: 'flex',
  alignItems: 'center',
};

function Navbar({ name }) {
  return (
    <AppBar
      position="fixed"
      sx={{ background: (theme) => theme.palette.navbarBackground.main }}
    >
      <Toolbar>
        <Container sx={containerStyles}>
          <Box sx={mobileLogoStyles}>
            <img src="/logo-mobile.svg" />
          </Box>
          <Box sx={desktopLogoStyles}>
            <img src="/logo.svg" />
          </Box>
          <Box sx={accountStyles}>
            <Typography variant="h6" mr="8px">
              Welcome, {name}!
            </Typography>
            <AccountCircleIcon />
          </Box>
        </Container>
      </Toolbar>
    </AppBar>
  );
}

export default Navbar;
