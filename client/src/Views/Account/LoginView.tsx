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

export type LoginViewProps = {
}

export default function LoginView(props: LoginViewProps) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const login = async () => {
    try {
      await fetch(GetLoginURL(), {
        method: 'POST',
        body: `userId=${username}`,
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
      })
      .then((response) => {
        if (response.ok) {
          response.text().then((data) => {
            const loginData = ParseLoginResponse(data);
            console.log(loginData);
            setLoginState(loginData);
            navigate('/')
          })
        } else {
          response.text().then((data) => {
            console.error(data);
            setErrorMessage(data);
          })
        }
      })
    } catch (e) {
      console.error(e);
    }
  }

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    await login();
    setIsLoading(false);
  }

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      justifyContent: 'center',
      alignItems: 'center',
    }}>
      <StyledForm
        sx={{
          display: 'grid',
          gridTemplateAreas: `
            "title"
            "username"
            "password"
            "submit"
            "error"
          `,
          textAlign: 'center',
          rowGap: '12px',
        }}
        onSubmit={onSubmit}>
        <Typography variant="h2">
          Log In
        </Typography>
        <TextField 
          onChange={(e) => setUsername(e.target.value)}
          label="userId"
          id="userId"
          type="text" />
        <TextField
          onChange={(e) => setPassword(e.target.value)}
          label="password"
          type="password"
          id="Password" />
        <Button type="submit">
          {
            isLoading ? (
              <CircularProgress />
            ) : (
            <Typography variant="body1">
              Log In
            </Typography>
            )
          }
        </Button>
        <Alert sx={{
          visibility: !!errorMessage ? 'visible' : 'hidden',
        }} severity="error">
          {errorMessage}
        </Alert>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
