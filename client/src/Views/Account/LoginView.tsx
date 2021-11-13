import React, { useState } from 'react';
import {
  Alert,
  Box, Button, CircularProgress, TextField, Typography,
} from '@mui/material';
import {
  styled
} from '@mui/material/styles';
import { GetLoginURL, ParseLoginResponse } from '../../Utils/BackendIntegration';
import { setLoginState } from '../../Utils/LoginState';
import { useNavigate } from 'react-router';

/**
 * The view used to describe the Login form for a user. Should only be visible the user is not
 * actively logged-in.
 * 
 * @returns The view.
 */
export default function LoginView() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  // TODO: Setup password auth.
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const _ = password;

  const login = async () => {
    try {
      await fetch(GetLoginURL(), {
        method: 'POST',
        body: `userId=${username}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      })
      .then((response) => {
        if (response.ok) {
          response.text().then((data) => {
            const loginData = ParseLoginResponse(data);
            console.log(loginData);
            setLoginState(loginData);
            navigate('/');
          });
        } else {
          response.text().then((data) => {
            console.error(data);
            setErrorMessage(data);
          });
        }
      });
    } catch (e) {
      console.error(e);
    }
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    await login();
    setIsLoading(false);
  };

  const onClickSignup = () => {
    navigate('/signup');
  };

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "."
        "form"
        "signup"
        "."
      `,
      justifyContent: 'center',
      alignItems: 'center',
      gridTemplateRows: '1fr min-content min-content 1fr',
      rowGap: '12px',
    }}>
      <StyledForm
        sx={{
          display: 'grid',
          gridTemplateAreas: `
            "title"
            "alert"
            "username"
            "password"
            "submit"
          `,
          textAlign: 'center',
          rowGap: '12px',
          gridArea: 'form',
        }}
        onSubmit={onSubmit}>
        <Typography variant="h2">
          Log In
        </Typography>
        <Alert sx={{
          visibility: !!errorMessage ? 'visible' : 'hidden',
        }} severity="error">
          {errorMessage}
        </Alert>
        <TextField 
          onChange={(e) => setUsername(e.target.value)}
          label="User ID"
          id="userId"
          type="text" />
        <TextField
          onChange={(e) => setPassword(e.target.value)}
          label="Password"
          type="password"
          id="Password" />
        <Button type="submit">
          {
            isLoading ? (
              <CircularProgress />
            ) : (
            <Typography variant="button">
              Log In
            </Typography>
            )
          }
        </Button>
      </StyledForm>
      <Button
        sx={{
          gridArea: 'signup',
        }}
        onClick={onClickSignup}>
        <Typography variant="button">
          Don't have an account? Sign up here.
        </Typography>
      </Button>
    </Box>
  );
}

const StyledForm = styled('form')``;
