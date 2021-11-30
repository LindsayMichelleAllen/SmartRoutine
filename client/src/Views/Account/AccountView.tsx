import {
  LoadingButton,
} from '@mui/lab';
import {
  Box,
  styled,
  Typography,
} from '@mui/material';
import React, {
  useEffect,
  useState,
} from 'react';
import AlertsBox from '../../Components/Containers/AlertsBox';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';
import {
  FetchRequest,
  ParseLoginResponse,
} from '../../Utils/BackendIntegration';
import {
  ValidateName,
  ValidateUsername,
} from '../../Utils/InputValidation';
import {
  useAuth,
} from '../../Utils/LoginState';

/**
 * The view used to describe a user's account and its specific details.
 *
 * @returns The view.
 */
export default function AccountView() {
  const [username, setUsername] = useState('');
  const [name, setName] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [genericError, setGenericError] = useState('');
  const [usernameValidation, setUsernameValidation] = useState('');
  const [nameValidation, setNameValidation] = useState('');
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
      const response = await FetchRequest('userUpdate', {
        name,
        username,
      });

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
      setGenericError(`${e}`);
    }
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    setIsLoading(true);
    await updateUser();
    setIsLoading(false);
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

  return (
    <Box sx={{
      display: 'grid',
      gridTemplateAreas: `
        "."
        "accountDetails"
        "."
      `,
      justifyContent: 'center',
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
        <AlertsBox
          successMessage={successMessage}
          errorMessage={genericError} />
        <StyledForm
          onSubmit={onSubmit}
          sx={{
            display: 'grid',
            gridArea: 'form',
            gridTemplateRows: 'auto',
            gridTemplateAreas: `
              "username"
              "name"
              "submit"
            `,
            rowGap: '12px',
          }}>
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
            labelText="Name"
            value={name}
            errorMessage={nameValidation}
            onValueChange={onNameChange} />
          <LoadingButton
            sx={{
              gridArea: 'submit',
            }}
            loading={isLoading}
            type="submit">
            <Typography variant="button">Update</Typography>
          </LoadingButton>
        </StyledForm>
      </Box>
    </Box>
  );
}

const StyledForm = styled('form')``;
