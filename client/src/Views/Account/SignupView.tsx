import {
  Alert,
  Box,
  TextField,
  Typography,
  styled,
} from '@mui/material';
import {
  FetchRequest,
  ParseLoginResponse,
} from '../../Utils/BackendIntegration';
import React, {
  useState,
} from 'react';
import { useAuth } from '../../Utils/LoginState';
import { ValidUserNameChars } from '../../Utils/InputValidation';
import { LoadingButton } from '@mui/lab';

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
  const authState = useAuth();

  const signIn = authState?.signIn;

  const signup = async () => {
    try {
      const response = await FetchRequest('userCreate', {
        name,
        username,
        password,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const loginData = ParseLoginResponse(text);

      if (signIn) {
        signIn(loginData);
        setSuccessMessage(`Success! The user account for ${loginData.Username} was created.`);
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
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
    } else if (!username.match(ValidUserNameChars)) {
      return 'Please only use letters and numbers in your username.';
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
        <LoadingButton loading={isLoading} sx={{ gridArea: 'submit' }} type="submit">
          <Typography variant="button">Sign Up</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
