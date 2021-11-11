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
import RoutinesView from '../Routines/RoutinesView';

export type AppProps = {
}

function App(props: AppProps) {
  const location = useLocation()

  return (
    <div>
      <MenuBar
        title={location.pathname}
      />
      <Routes>
        <Route
          path="/"
          element={<Navigate replace to="/account"/>} />
        <Route
          path="/account"
          element={<AccountView />} />
        <Route
          path="/routines"
          element={<RoutinesView />} />
      </Routes>
    </div>
  );
}

export default function AppContainer() {
  return (
    <BrowserRouter>
      <App />
    </BrowserRouter>
  )
}
