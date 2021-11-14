import {
  Divider,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Paper,
  Toolbar,
  Typography,
} from '@mui/material';
import React, { useMemo } from 'react';
import { useNavigate } from 'react-router';
import { useAuth } from '../../Utils/LoginState';
import LoginIcon from '@mui/icons-material/Login';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import AltRouteIcon from '@mui/icons-material/AltRoute';
import {
  ACCOUNT_URL,
  LOGIN_URL,
  ROUTINES_URL,
  LOGOUT_URL,
} from '../../Utils/CommonRouting';

/**
 * A generic type used to describe a singular navigation item in the nav view.
 */
type NavItem = {
  /**
   * The icon associated with this element. There's not very strong typing here, so be certain to
   * include an icon from the @mui/icons-material library.
   */
  icon: JSX.Element;

  /**
   * The label that will be rendered next to the nav item.
   */
  label: string;

  /**
   * The route associated with the given nav item.
   */
  route: string;
};

/**
 * A list of static navigation items to be included in the nav pane. These items require a user to
 * be logged-in in order for them to be displayed. This isn't route protection, but it helps a user
 * to understand what is available.
 */
const protectedNavItems: NavItem[] = [
  {
    icon: (<AccountCircleIcon />),
    label: 'Account',
    route: ACCOUNT_URL,
  },
  {
    icon: (<AltRouteIcon />),
    label: 'Routines',
    route: ROUTINES_URL,
  }
];


/**
 * A list of static navigation items to be included in the nav pane. These items do not require a
 * user to be logged-in in order for them to be displayed.
 */
const unprotectedNavItems: NavItem[] = [
];

/**
 * The props for the {@link NavDrawer} element.
 */
export type NavDrawerProps = {
  /**
   * The drawer width size, in pixels. This should match what is expected from the parent view.
   */
  drawerWidthPixels: number;

  /**
   * A callback event for when the drawer attempts to navigate somewhere. Useful for mobile views
   * where the user wants the pane closed after navigating.
   */
  onNavigate?: () => void;
}

/**
 * A component used to provide a set of navigation options for the user.
 * 
 * @param props See {@link NavDrawerProps}.
 * @returns The component.
 */
export default function NavDrawer(props: NavDrawerProps) {
  const {
    drawerWidthPixels,
    onNavigate,
  } = props;

  const navigate = useNavigate();
  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const handleNavigate = (route: string) => {
    navigate(route);
    if (onNavigate) {
      onNavigate();
    }
  };

  const navListItems = useMemo(() => {
    const logOption: NavItem = loginDetails === undefined
    ? {
      icon: (<LoginIcon />),
      label: 'Login Page',
      route: LOGIN_URL,
    }
    : {
      icon: (<LoginIcon />),
      label: 'Logout Page',
      route: LOGOUT_URL,
    };

    const navOptions = [
      ...unprotectedNavItems,
      ...(loginDetails !== undefined ? protectedNavItems : []),
      logOption,
    ];

    const mappedItems = navOptions.map((n) => (
      <ListItem key={n.label} button onClick={() => handleNavigate(n.route)}>
        <ListItemAvatar>
          {n.icon}
        </ListItemAvatar>
        <ListItemText primary={n.label} />
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
