import React from 'react';
import { useLoginState } from '../../Utils/LoginState';
import LoginView from './LoginView';
import LogoutView from './LogoutView';

/**
 * A view wrapper used to render either the login or the logout view based on the user's active
 * login session.
 * 
 * @returns The view.
 */
export default function LoginRouter() {
  const loginState = useLoginState();

  // TODO: Make this more of an authentication than a simple state check.
  const component: JSX.Element = loginState && loginState.userid ? (
    <LogoutView />
  ) : (
    <LoginView />
  );

  return component;
}
