import { Box, CircularProgress } from '@mui/material';
import React, { useMemo } from 'react';
import { Navigate } from 'react-router';
import { useAuth } from '../../Utils/LoginState';

/**
 * Props for the {@link PrivateRoute} element.
 */
export type PrivateRouteProps = {
  /**
   * The element to render, if the user's authorization matches the expected state.
   */
  authElement?: JSX.Element;

  /**
   * The fallback URL to re-direct the user to if their authenticated state is not what was
   * expected.
   */
  fallbackUrl?: string;

  /**
   * If true, re-direct the user if they 'are' authenticated. Otherwise, re-direct the user if they
   * 'are not' authenticated.
   */
  invertPrivacy?: boolean;
};

/**
 * An element used to re-direct the user based on their current authentication.
 * 
 * @param props See {@link PrivateRouteProps}.
 * @returns The element.
 */
export default function PrivateRoute(props: PrivateRouteProps) {
  const {
    authElement,
    fallbackUrl,
    invertPrivacy,
  } = props;

  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const doReturnElement = !!invertPrivacy !== !!loginDetails;
  const loadingElement = useMemo(() => (
    <Box>
      <CircularProgress />
    </Box>
  ), []);

  if (!authState.attemptedToGetState) {
    return loadingElement;
  }
  if (!doReturnElement) {
    return (<Navigate to={fallbackUrl} />);
  }

  return authElement;
}
