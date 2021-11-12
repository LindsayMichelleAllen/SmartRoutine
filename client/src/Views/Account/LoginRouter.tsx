import { createTheme, ThemeProvider } from '@mui/material/styles';
import React from 'react';
import { useLoginState } from '../../Utils/LoginState';
import LoginView from './LoginView';
import LogoutView from './LogoutView';

export type LoginRouterProps = {
}

export default function LoginRouter(props: LoginRouterProps) {
  const loginState = useLoginState();

  // TODO: Make this more of an authentication than a simple state check.
  const component: JSX.Element = loginState && loginState.userid ? (
    <LogoutView />
  ) : (
    <LoginView />
  )

  return component;
}
