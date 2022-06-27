import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted';
import SettingsIcon from '@mui/icons-material/Settings';
import List from '@mui/material/List';
import Divider from '@mui/material/Divider';
import { Link } from 'react-router-dom';
import { useRecoilValue } from 'recoil';
import { PageState } from '../../pages/store';

export default function NavList() {
  const pageState = useRecoilValue(PageState);

  return (
    <List component="nav">
      <>
        <ListItemButton
          component={Link}
          to="/"
          selected={pageState.index === 'home'}
        >
          <ListItemIcon>
            <DashboardIcon />
          </ListItemIcon>
          <ListItemText primary="Home" />
        </ListItemButton>
        <ListItemButton
          component={Link}
          to="/devices"
          selected={pageState.index === 'devices'}
        >
          <ListItemIcon>
            <FormatListBulletedIcon />
          </ListItemIcon>
          <ListItemText primary="Devices" />
        </ListItemButton>
      </>
      <Divider sx={{ my: 1 }} />
      <>
        <ListSubheader component="div" inset>
          Utilities
        </ListSubheader>
        <ListItemButton
          component={Link}
          to="/settings"
          selected={pageState.index === 'settings'}
        >
          <ListItemIcon>
            <SettingsIcon />
          </ListItemIcon>
          <ListItemText primary="Settings" />
        </ListItemButton>
      </>
    </List>
  );
}
