import { Divider, List, ListItem, ListItemAvatar, ListItemText, Paper, Toolbar, Typography } from '@mui/material';
import React, { useMemo } from 'react';
import { useNavigate } from 'react-router';
import { useAuth } from '../../Utils/LoginState';
import LoginIcon from '@mui/icons-material/Login';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import AltRouteIcon from '@mui/icons-material/AltRoute';

type NavItem = {
  icon: JSX.Element;
  name: string;
  route: string;
};

const navItems: NavItem[] = [
  {
    icon: (<AccountCircleIcon />),
    name: 'Account',
    route: '/account',
  },
  {
    icon: (<AltRouteIcon />),
    name: 'Routines',
    route: '/routines',
  }
];

export type NavDrawerProps = {
  drawerWidthPixels: number;
  onNavigate?: () => void;
}

/**
 * @param props
 */
export default function NavDrawer(props: NavDrawerProps) {
  const {
    drawerWidthPixels,
    onNavigate,
  } = props;

  const navigate = useNavigate();
  const { loginDetails } = useAuth();

  const handleNavigate = (route: string) => {
    navigate(route);
    if (onNavigate) {
      onNavigate();
    }
  };

  const navListItems = useMemo(() => {
    const logOption = loginDetails === undefined
    ? {
      icon: (<LoginIcon />),
      name: 'Login Page',
      route: '/login',
    }
    : {
      icon: (<LoginIcon />),
      name: 'Logout Page',
      route: '/logout',
    };

    const mappedItems = [...navItems, logOption].map((n) => (
      <ListItem key={n.name} button onClick={() => handleNavigate(n.route)}>
        <ListItemAvatar>
          {n.icon}
        </ListItemAvatar>
        <ListItemText primary={n.name} />
      </ListItem>
    ));

    return [...mappedItems];
  }, [loginDetails]);

  return (
    <Paper sx={{ width: `${drawerWidthPixels}px` }}>
      <Toolbar>
        <Typography variant="h5">
          SmartRoutine
        </Typography>
      </Toolbar>
      <Divider />
      <List>
        {navListItems}
      </List>
    </Paper>
  );
}
