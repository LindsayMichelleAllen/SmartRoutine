import React from 'react';
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  useLocation,
} from 'react-router-dom';
import MenuBar from '../../Components/MenuBar/MenuBar';
import AccountView from '../Account/AccountView';
import LoginRouter from '../Account/LoginRouter';
import LoginView from '../Account/LoginView';
import SignupView from '../Account/SignupView';
import RoutinesView from '../Routines/RoutinesView';
import { Box, createTheme, Paper, ThemeProvider } from '@mui/material';
import PrivateRoute from '../../Components/Routing/PrivateRoute';

const theme = createTheme({
  components: {
    MuiButton: {
      defaultProps: {
        variant: 'contained',
      },
    },
  },
});

/**
 * The main entry point for the application.
 * 
 * @returns The view.
 */
function App() {
  const location = useLocation();

  return (
    <Box sx={{
      display: 'grid',
      height: '100%',
      width: '100%',
      gridTemplateAreas: `
        "menubar"
        "body"
      `,
      gridTemplateRows: 'min-content auto'
    }}>
      <MenuBar
        title={location.pathname}
      />
      <Paper>
        <Routes>
          <Route
            path="/"
            element={<Navigate replace to="/login" />} />
          <Route
            path="/login"
            element={<LoginRouter />} />
          <Route
            path="/signup"
            element={<SignupView />} />
          <Route
            path="/account"
            element={<AccountView />} />
          <Route
            path="/routines"
            element={<RoutinesView />} />
          <Route
            path="/login"
            element={<LoginView />} />
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
  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <App />
      </ThemeProvider>
    </BrowserRouter>
  );
}
