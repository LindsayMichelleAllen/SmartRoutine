import { Box, Button, Typography } from "@mui/material";
import React from "react";
import { setLoginState, useLoginState } from "../../Utils/LoginState";
import { useNavigate } from 'react-router';

export type LogoutViewProps = {
}

export default function LogoutView(props: LogoutViewProps) {
  const loginDetails = useLoginState();
  const navigate = useNavigate();

  const handleClick = () => {
    setLoginState(undefined);
    navigate('/');
  }

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
          <Typography variant="body1">
            Log Out
          </Typography>
        </Button>
      </Box>
    </Box>
  );
}
