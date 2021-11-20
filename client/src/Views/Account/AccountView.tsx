import {
  Alert,
  Box,
  Button,
  CircularProgress,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import React, { useEffect, useState } from 'react';
import {
  GetFetchRequest,
  GetModifyUserURL,
  ParseLoginResponse,
} from '../../Utils/BackendIntegration';
import { ValidUserNameChars } from '../../Utils/InputValidation';
import { useAuth } from '../../Utils/LoginState';

/**
 * The view used to describe a user's account and its specific details.
 *
 * @returns The view.
 */
export default function AccountView() {
  const [username, setUsername] = useState('');
  const [name, setName] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const authState = useAuth();
  const loginDetails = authState?.loginDetails;
  const signIn = authState?.signIn;

  useEffect(() => {
    if (loginDetails) {
      setUsername(loginDetails.Username);
      setName(loginDetails.Name);
    }
  }, [loginDetails]);

  const updateUser = async () => {
    try {
      const response = await fetch(
        GetModifyUserURL(),
        GetFetchRequest({
          username,
          name,
        }),
      );

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const loginData = ParseLoginResponse(text);
      if (signIn) {
        signIn(loginData);
        setUsername(loginData.Username);
        setName(loginData.Name);
        setSuccessMessage(`
          Success! Your account has been updated.
        `);
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }
  };

  const validateInput = (): string | undefined => {
    if (!username.match(ValidUserNameChars)) {
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
      await updateUser();
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
          "title"
          "form"
        `,
        rowGap: '12px',
        alignItems: 'center',
      }}>
        <Typography sx={{ gridArea: 'title' }} variant="h3">
          Account Details
        </Typography>
        <StyledForm
          onSubmit={onSubmit}
          sx={{
            display: 'grid',
            gridArea: 'form',
            gridTemplateRows: 'auto',
            gridTemplateAreas: `
              "error"
              "success"
              "username_label"
              "username"
              "name_label"
              "name"
              "submit"
            `,
            rowGap: '12px',
          }}>
          <Alert severity="success"
            sx={{
              gridArea: 'success',
              visibility: !!successMessage ? 'visible' : 'hidden'
            }}>
            {successMessage}
          </Alert>
          <Alert severity="error" sx={{
            gridArea: 'error',
            visibility: !!errorMessage ? 'visible' : 'hidden',
          }}>
            {errorMessage}
          </Alert>
          <Typography sx={{ gridArea: 'username_label' }} variant="h6">
            Username
          </Typography>
          <TextField
            sx={{ gridArea: 'username' }}
            onChange={(e) => setUsername(e.target.value)}
            value={username}
            label="Username"
            id="username"
            type="text" />
          <Typography sx={{ gridArea: 'name_label' }} variant="h6">
            Name
          </Typography>
          <TextField
            sx={{ gridArea: 'name' }}
            onChange={(e) => setName(e.target.value)}
            value={name}
            label="Name"
            id="name"
            type="text" />
          <Button sx={{ gridArea: 'submit' }} type="submit">
            {
              isLoading
                ? (<CircularProgress />)
                : (<Typography variant="button">Update</Typography>)
            }
          </Button>
        </StyledForm>
      </Box>
    </Box>
  );
}

const StyledForm = styled('form')``;
