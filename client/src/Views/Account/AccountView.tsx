import { Box, Typography } from '@mui/material';
import React from 'react';
import { useAuth } from '../../Utils/LoginState';

/**
 * The view used to describe a user's account and its specific details.
 *
 * @returns The view.
 */
export default function AccountView() {
  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "."
        "accountDetails"
        "."
      `,
      justifyContent: 'center',
      alignItems: 'center',
    }}>
      <Box sx={{
        gridArea: 'accountDetails',
        display: 'grid',
        gridTemplateAreas: `
          "title title"
          "usernameTitle username"
          "nameTitle name"
          "userIdTitle userId"
        `,
        rowGap: '12px',
        alignItems: 'center',
        columnGap: '48px',
        justifyContent: 'space-between',
      }}>
        <Typography sx={{ gridArea: 'title' }} variant="h2">
          Account Details
        </Typography>
        <Typography sx={{ gridArea: 'usernameTitle' }} variant="h6">
          Username
        </Typography>
        <Typography sx={{ gridArea: 'username' }} variant="body1">
          {loginDetails?.username ?? ''}
        </Typography>
        <Typography sx={{ gridArea: 'nameTitle' }} variant="h6">
          Name
        </Typography>
        <Typography sx={{ gridArea: 'name' }} variant="body1">
          {loginDetails?.name ?? ''}
        </Typography>
        <Typography sx={{ gridArea: 'userIdTitle', }} variant="h6">
          User ID
        </Typography>
        <Typography sx={{ gridArea: 'userId' }} variant="body1">
          {loginDetails?.userid ?? ''}
        </Typography>
      </Box>
    </Box>
  );
}
