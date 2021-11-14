import React from 'react';
import { Navigate } from 'react-router';
import { useAuth } from '../../Utils/LoginState';

export type PrivateRouteProps = {
  authElement?: JSX.Element;
  fallbackUrl?: string;
  invertPrivacy?: boolean;
};

/**
 * @param props
 */
export default function PrivateRoute(props: PrivateRouteProps) {
  const {
    authElement,
    fallbackUrl,
    invertPrivacy,
  } = props;

  const { loginDetails } = useAuth();
  const doReturnElement = !!invertPrivacy !== !!loginDetails;
  return doReturnElement ? authElement : (<Navigate to={fallbackUrl} />);
}
