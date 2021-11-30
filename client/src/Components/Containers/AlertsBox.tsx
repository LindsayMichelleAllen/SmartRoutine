import {
  Box,
  BoxProps,
  Alert,
} from '@mui/material';
import React from 'react';

/**
 * See {@link AlertsBox}.
 */
export type AlertsBoxProps = BoxProps & {
  /**
   * The success message to display in an alert. If this is undefined/empty, the alert will not
   * take up vertical space.
   */
  successMessage?: string;

  /**
   * The error message to display in an alert. If this is undefined/empty, the alert will not take
   * up vertical space.
   */
  errorMessage?: string;
}

/**
 * A component used to wrap success and error alerts.
 * 
 * @param props See {@link AlertsBoxProps}.
 * @returns The component.
 */
export default function AlertsBox(props: AlertsBoxProps) {
  const {
    successMessage,
    errorMessage,
  } = props;

  const successAlert = (
    <Alert
      sx={{
        visibility: !!successMessage ? 'visible' : 'hidden',
        height: !!successMessage ? 'auto' : 0,
      }}
      severity="success"
      variant="outlined" >
      {successMessage}
    </Alert>
  );

  const errorAlert = (
    <Alert
      sx={{
        visibility: !!errorMessage ? 'visible' : 'hidden',
        height: !!errorMessage ? 'auto' : 0,
      }}
      severity="error"
      variant="outlined" >
      {errorMessage}
    </Alert>
  );

  return (
    <Box sx={{
      display: 'grid',
      gridAutoRows: 'auto',
      justifyContent: 'center',
    }}>
      {errorAlert}
      {successAlert}
    </Box>
  );
}
