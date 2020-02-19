import * as React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import { Link } from 'react-router-dom'
import { Menu, MenuItem, Button, Grid, AppBar } from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

const PageHeader = () => {
  const [anchorEl, setAnchorEl] = React.useState(null)
  const classes = useStyles({})

  const handleClick = event => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <header className={classes.root}>
      <AppBar position="static">
        <Grid container
          direction="row"
          justify="flex-end"
          alignItems="center">
          <Button aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
            <MenuIcon></MenuIcon>
          </Button>
          <Menu
            id="simple-menu"
            anchorEl={anchorEl}
            keepMounted
            open={Boolean(anchorEl)}
            onClose={handleClose}
          >
            <MenuItem
              onClick={handleClose}
              component={Link}
              to="/">Home</MenuItem>
            <MenuItem
              onClick={handleClose}
              component={Link}
              to="/about">About</MenuItem>
            <MenuItem
              onClick={handleClose}
              component={Link}
              to="/device-management">Device Management</MenuItem>
            <MenuItem
              onClick={handleClose}
              component={Link}
              to="/dashboard/edit">Dashboard Edit</MenuItem>
          </Menu>
        </Grid>
      </AppBar>

    </header>
  )
}


export default PageHeader
