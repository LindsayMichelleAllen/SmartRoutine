import { createTheme, ThemeProvider } from '@mui/material/styles';
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
import {
  styled
} from '@mui/material/styles';
import { Box, Paper } from '@mui/material';

export type AppProps = {
}

function App(props: AppProps) {
  const location = useLocation()

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
      <div>
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
      </div>
    </Box>
  );
}

export default function AppContainer() {
  return (
    <BrowserRouter>
      <App />
    </BrowserRouter>
  )
}
