import {
  Alert,
  Box,
  Button,
  CircularProgress,
  TextField,
  Typography,
  styled,
} from '@mui/material';
import {
  GetSignupURL, ParseLoginResponse,
} from '../../Utils/BackendIntegration';
import React, {
  useState,
} from 'react';
import { useAuth } from '../../Utils/LoginState';

const validUserChars = /^[0-9a-zA-Z]+$/;

/**
 * The view used to provide the user with a means to create a new account.
 * 
 * @returns The view.
 */
export default function SignupView() {
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const [username, setUsername] = useState('');
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { signIn } = useAuth();

  const signup = async () => {
    try {
      const response = await fetch(GetSignupURL(), {
        method: 'POST',
        body: `username=${username}&name=${name}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const loginData = ParseLoginResponse(text);

      signIn(loginData);
      setSuccessMessage(`
        Success! The user account for ${loginData.username} was created.\n
        Your user ID is "${loginData.userid}".\n
        Use this ID to log in again.`);
    } catch (e) {
      console.error(e);
    } finally {
      setUsername('');
      setName('');
      setPassword('');
      setConfirmPassword('');
    }
  };

  const validateInput = (): string | undefined => {
    if (password !== confirmPassword) {
      return 'Passwords do not match.';
    } else if (!username.match(validUserChars)) {
      return 'Please only use letters and numbers in your username.';
    } else if (!name.match(validUserChars)) {
      return 'Please only use letters and numbers in your name.';
    }

    return undefined;
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsLoading(true);

    const validationError = validateInput();
    if (validationError !== undefined) {
      setErrorMessage(validationError);
    } else {
      await signup();
    }

    setIsLoading(false);
  };

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "."
        "form"
        "."
      `,
      justifyContent: 'center',
      alignItems: 'center',
      gridTemplateRows: '1fr min-content 1fr',
    }}>
      <StyledForm
        sx={{
          display: 'grid',
          gridTemplateAreas: `
            "title"
            "error"
            "success"
            "username"
            "name"
            "password"
            "confirmpassword"
            "submit"
          `,
          textAlign: 'center',
          rowGap: '12px',
          gridArea: 'form',
        }}
        onSubmit={onSubmit} >
        <Typography variant="h2">
          Sign Up
        </Typography>
        <Alert sx={{
          visibility: !!errorMessage ? 'visible' : 'hidden',
        }} severity="error">
          {errorMessage}
        </Alert>
        <Alert sx={{
          visibility: !!successMessage ? 'visible' : 'hidden',
        }} severity="success">
          {successMessage}
        </Alert>
        <TextField
          onChange={(e) => setUsername(e.target.value)}
          value={username}
          label="Username"
          id="username"
          type="text" />
        <TextField
          onChange={(e) => setName(e.target.value)}
          value={name}
          label="Your Name"
          id="name"
          type="text" />
        <TextField
          onChange={(e) => setPassword(e.target.value)}
          value={password}
          label="Password"
          id="password"
          type="password" />
        <TextField
          onChange={(e) => setConfirmPassword(e.target.value)}
          value={confirmPassword}
          label="Confirm Password"
          id="confirmpassword"
          type="password" />
        <Button type="submit">
          {
            isLoading ? (
              <CircularProgress />
            ) : (
              <Typography variant="button">
                Sign Up
              </Typography>
            )
          }
        </Button>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
