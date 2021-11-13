import React from 'react';
import { Navigate, Route, RouteProps } from 'react-router';
import { useLoginState } from '../../Utils/LoginState';

export type PrivateRouteProps = {
  element?: JSX.Element;
};

export default function PrivateRoute(props: PrivateRouteProps) {
  const {
    element,
  } = props;

  const loginState = useLoginState();
  return !!loginState ? element : (<Navigate to="/login" />);
}
