import React, { useState } from 'react';
import {
  Alert,
  Box,
  Button,
  CircularProgress,
  TextField,
  Typography,
} from '@mui/material';
import {
  styled
} from '@mui/material/styles';
import {
  FetchRequest,
  ParseLoginResponse,
} from '../../Utils/BackendIntegration';
import { useNavigate } from 'react-router';
import { useAuth } from '../../Utils/LoginState';

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

  const authState = useAuth();
  const signIn = authState?.signIn;


  // TODO: Setup password auth.
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const _ = password;

  const login = async () => {
    try {
      const response = await FetchRequest('userLogin', {
        username,
        password,
      });

      const text = await response.text();
      if (response.ok) {
        const loginData = ParseLoginResponse(text);
        if (signIn) {
          signIn(loginData);
        }
      } else {
        throw text;
      }
    } catch (e) {
      throw e;
    }
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      await login();
      navigate('/');
    } catch (e) {
      // Only set loading to false if the login failed. If we try to set it on a success, that gets
      // called after navigate which leads to the 'memory leak' React error.
      console.error(e);
      setErrorMessage(`${e}`);
      setIsLoading(false);
    }
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
