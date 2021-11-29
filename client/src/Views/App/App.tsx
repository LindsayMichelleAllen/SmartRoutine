import React, {
  useEffect,
  useMemo,
  useState,
} from 'react';
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  useLocation,
} from 'react-router-dom';
import MenuBar from '../../Components/MenuBar/MenuBar';
import AccountView from '../Account/AccountView';
import LoginView from '../Account/LoginView';
import SignupView from '../Account/SignupView';
import RoutinesView from '../Routines/RoutinesView';
import {
  Box,
  createTheme,
  Divider,
  Drawer,
  PaletteMode,
  Paper,
  ThemeProvider,
  useTheme,
} from '@mui/material';
import PrivateRoute from '../../Components/Routing/PrivateRoute';
import { AuthProvider } from '../../Utils/LoginState';
import {
  ColorProvider,
  fetchColorPreference,
  storeColorPreference,
} from '../../Utils/ColorContext';
import NavDrawer from '../../Components/NavDrawer/NavDrawer';
import LogoutView from '../Account/LogoutView';
import {
  ACCOUNT_URL,
  ADD_DEVICE_TO_ROUTINE_URL,
  ADD_DEVICE_URL,
  ADD_ROUTINE_URL,
  DEVICES_URL,
  EDIT_DEVICE_URL,
  EDIT_ROUTINE_URL,
  LOGIN_URL,
  LOGOUT_URL,
  ROUTINES_URL,
  SIGNUP_URL,
  VIEW_ROUTINE_URL,
} from '../../Utils/CommonRouting';
import AddRoutineView from '../Routines/AddRoutineView';
import RoutineView from '../Routines/SingleRoutineView';
import DevicesView from '../Devices/DevicesView';
import {
  LocalizationProvider,
} from '@mui/lab';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import EditRoutineView from '../Routines/EditRoutineView';
import EditDeviceView from '../Devices/EditDeviceView';
import AddDeviceView from '../Devices/AddDeviceView';
import AddConfigurationView from '../Routines/AddConfigurationView';

const navDrawerWidth = 240;

/**
 * The main entry point for the application.
 * 
 * @returns The view.
 */
function App() {
  const location = useLocation();
  const theme = useTheme();
  const [mobileOpen, setMobileOpen] = useState(false);

  const formattedLocation = useMemo(() => {
    let formattedRoute = '';

    if (location.pathname.includes('/')) {
      formattedRoute = location.pathname.split('/')[1];
    }

    if (formattedRoute.length > 1) {
      formattedRoute = formattedRoute.slice(0, 1).toUpperCase() + formattedRoute.slice(1);
    }

    return formattedRoute;
  }, [location]);

  return (
    <Box sx={{
      display: 'grid',
      height: '100%',
      width: '100%',
      gridTemplateAreas: `
        "nav menubar"
        "nav divider"
        "nav body"
      `,
      gridTemplateRows: 'min-content min-content 1fr',
      gridTemplateColumns: {
        xs: '0px auto',
        sm: `${navDrawerWidth}px auto`,
      }
    }}
    >
      <Box sx={{gridArea: 'nav'}} component="nav" >
        <Drawer
          variant="temporary"
          open={mobileOpen}
          onClose={() => setMobileOpen(false)}
          ModalProps={{
            keepMounted: true, // https://mui.com/components/drawers/ Better Perf on Mobile.
          }}
          sx = {{ display: { xs: 'block', sm: 'none' } }}>
          <NavDrawer onNavigate={() => setMobileOpen(false)} drawerWidthPixels={navDrawerWidth} />
        </Drawer>
        <Drawer
          variant="permanent"
          open
          sx={{ display: { xs: 'none', sm: 'block' } }}>
          <NavDrawer drawerWidthPixels={navDrawerWidth} />
        </Drawer>
      </Box>
      <MenuBar
        sx={{ gridArea: 'menubar' }}
        title={formattedLocation}
        handleClickMenu={() => setMobileOpen(true)}
      />
      <Divider sx={{ gridArea: 'divider', borderColor: theme.palette.mode === 'dark' ? 'black' : 'white' }} />
      <Paper square sx={{ overflowY: 'auto' }} >
        <Routes>
          <Route
            path="/"
            element={<Navigate replace to={LOGIN_URL} />} />
          <Route
            path={LOGOUT_URL}
            element={<PrivateRoute authElement={<LogoutView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={SIGNUP_URL}
            element={<PrivateRoute authElement={<SignupView />} fallbackUrl={LOGOUT_URL} invertPrivacy/>} />
          <Route
            path={ACCOUNT_URL}
            element={<PrivateRoute authElement={<AccountView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={ROUTINES_URL}
            element={<PrivateRoute authElement={<RoutinesView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={DEVICES_URL}
            element={<PrivateRoute authElement={<DevicesView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={ADD_ROUTINE_URL}
            element={<PrivateRoute authElement={<AddRoutineView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={ADD_DEVICE_URL}
            element={<PrivateRoute authElement={<AddDeviceView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={EDIT_ROUTINE_URL}
            element={<PrivateRoute authElement={<EditRoutineView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={EDIT_DEVICE_URL}
            element={<PrivateRoute authElement={<EditDeviceView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={VIEW_ROUTINE_URL}
            element={<PrivateRoute authElement={<RoutineView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={ADD_DEVICE_TO_ROUTINE_URL}
            element={<PrivateRoute authElement={<AddConfigurationView />} fallbackUrl={LOGIN_URL} />} />
          <Route
            path={LOGIN_URL}
            element={<PrivateRoute authElement={<LoginView />} fallbackUrl={LOGOUT_URL} invertPrivacy/>} />
        </Routes>
      </Paper>
    </Box>
  );
}

/**
 * A wrapper for the default app container. Provides some high-level wrapper elements before
 * providing the {@link App} view.
 * 
 * @returns The wrapper.
 */
export default function AppContainer() {
  const [colorMode, setColorMode] = useState<PaletteMode>('dark');

  const theme = useMemo(() => createTheme({
    palette: {
      primary: {
        main: '#663399',
      },
      mode: colorMode,
    },
    components: {
      MuiFab: {
        styleOverrides: {
          root: {
            position: 'absolute',
            bottom: '24px',
            right: '24px',
          }
        },
        defaultProps: {
          color: 'primary',
        },
      },
      MuiCardActions: {
        styleOverrides: {
          root: {
            justifyContent: 'end',
          }
        },
      },
      MuiCard: {
        defaultProps: {
          variant: 'outlined'
        },
        styleOverrides: {
          root: {
            minHeight: '200px',
            minWidth: '320px',
          }
        }
      },
      MuiButton: {
        defaultProps: {
          variant: 'contained',
        },
      },
      MuiAppBar: {
        defaultProps: {
          enableColorOnDark: true,
        }
      },
    },
  }), [colorMode]);

  const updateColorPreference = (preferredColorMode: PaletteMode) => {
    storeColorPreference(preferredColorMode);
    setColorMode(preferredColorMode);
  };

  useEffect(() => {
    const colorPreference = fetchColorPreference();
    setColorMode(colorPreference);
  }, []);

  return (
    <BrowserRouter>
      <ColorProvider colorState={{colorMode, setColorMode: updateColorPreference}}>
        <ThemeProvider theme={theme}>
          <LocalizationProvider dateAdapter={AdapterDateFns}>
            <AuthProvider>
              <App />
            </AuthProvider>
          </LocalizationProvider>
        </ThemeProvider>
      </ColorProvider>
    </BrowserRouter>
  );
}
