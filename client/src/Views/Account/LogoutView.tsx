import {
  Box,
  Button,
  Typography,
} from '@mui/material';
import React from 'react';
import { useAuth } from '../../Utils/LoginState';
import { useNavigate } from 'react-router';

/**
 * The view used to provide the user with a logout option.
 * 
 * @returns The view.
 */
export default function LogoutView() {
  const authState = useAuth();
  const navigate = useNavigate();

  const loginDetails = authState?.loginDetails;
  const signOut = authState?.signOut;

  const handleClick = () => {
    if (signOut) {
      signOut();
    }
    navigate('/');
  };

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      justifyContent: 'center',
      alignItems: 'center',
    }}>
      <Box sx={{
        display: 'grid',
        gridTemplateAreas: `
          "title"
          "username"
          "submit"
        `,
        textAlign: 'center',
        rowGap: '12px',
      }}>
        <Typography variant="h2">
          Logged In
        </Typography>
        <Typography>
          {`Currently logged in as ${loginDetails?.username ?? ''}`}
        </Typography>
        <Button onClick={() => handleClick()}>
          <Typography variant="button">
            Log Out
          </Typography>
        </Button>
      </Box>
    </Box>
  );
}
