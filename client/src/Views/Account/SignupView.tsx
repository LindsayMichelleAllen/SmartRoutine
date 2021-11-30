import {
  Box,
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
import {
  useAuth,
} from '../../Utils/LoginState';
import {
  ValidateName,
  ValidatePassword,
  ValidateUsername,
} from '../../Utils/InputValidation';
import {
  LoadingButton,
} from '@mui/lab';
import AlertsBox from '../../Components/Containers/AlertsBox';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';

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

  const [usernameValidation, setUsernameValidation] = useState('');
  const [nameValidation, setNameValidation] = useState('');
  const [passwordValidation, setPasswordValidation] = useState('');
  const [confirmPasswordValidation, setConfirmPasswordValidation] = useState('');

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

  const validateInput = (): string | undefined => {
    if (password !== confirmPassword) {
      return 'Passwords do not match.';
    }

    return undefined;
  };

  const onUsernameChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setUsername(input);
    const validationError = ValidateUsername(input);
    if (!!validationError) {
      setUsernameValidation(validationError);
    } else {
      setUsernameValidation('');
    }
  };

  const onNameChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setName(input);
    const validationError = ValidateName(input);
    if (!!validationError) {
      setNameValidation(validationError);
    } else {
      setNameValidation('');
    }
  };

  const onPasswordChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setPassword(input);
    const validationError = ValidatePassword(input);
    if (!!validationError) {
      setPasswordValidation(validationError);
    } else {
      setPasswordValidation('');
    }
  };

  const onConfirmPasswordChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setConfirmPassword(input);
    const validationError = ValidatePassword(input);
    if (!!validationError) {
      setConfirmPasswordValidation(validationError);
    } else {
      setConfirmPasswordValidation('');
    }
  };

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "title"
        "form"
        "."
      `,
      justifyContent: 'center',
      alignItems: 'center',
      gridTemplateRows: '1fr min-content 1fr',
    }}>
      <Typography variant="h2">
        Sign Up
      </Typography>
      <AlertsBox
        errorMessage={errorMessage}
        successMessage={successMessage} />
      <StyledForm
        sx={{
          display: 'grid',
          gridTemplateAreas: `
            "username"
            "name"
            "password"
            "confirm-password"
            "submit"
          `,
          textAlign: 'center',
          rowGap: '12px',
          gridArea: 'form',
        }}
        onSubmit={onSubmit} >
        <ValidatedInput
          sx={{
            gridArea: 'username',
          }}
          labelId="username"
          labelText="Username"
          value={username}
          errorMessage={usernameValidation}
          onValueChange={onUsernameChange} />
        <ValidatedInput
          sx={{
            gridArea: 'name',
          }}
          labelId="name"
          labelText="Your Name"
          value={name}
          errorMessage={nameValidation}
          onValueChange={onNameChange} />
        <ValidatedInput
          sx={{
            gridArea: 'password',
          }}
          labelId="password"
          labelText="Password"
          value={password}
          errorMessage={passwordValidation}
          onValueChange={onPasswordChange} />
        <ValidatedInput
          sx={{
            gridArea: 'confirm-password',
          }}
          labelId="confirm-password"
          labelText="Confirm Password"
          value={confirmPassword}
          errorMessage={confirmPasswordValidation}
          onValueChange={onConfirmPasswordChange} />
        <LoadingButton loading={isLoading} sx={{ gridArea: 'submit' }} type="submit">
          <Typography variant="button">Sign Up</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
