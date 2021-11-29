import { LoadingButton } from '@mui/lab';
import {
  Alert,
  Box,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import React, { useState } from 'react';
import { useNavigate } from 'react-router';
import {
  FetchRequest,
  ParseDevice,
} from '../../Utils/BackendIntegration';
import { DEVICES_URL } from '../../Utils/CommonRouting';
import { ValidDeviceNameChars } from '../../Utils/InputValidation';
import { useAuth } from '../../Utils/LoginState';

/**
 * AddDeviceView provides the means for a user to create a new device.
 * 
 * @returns The view.
 */
export default function AddDeviceView() {
  const [name, setName] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();

  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const addDevice = async () => {
    try {
      const response = await FetchRequest('deviceCreate', {
        name,
        userid: loginDetails.Username,
      });

      const text = await response.text();

      if (response.ok) {
        const deviceData = ParseDevice(text);
        if (deviceData !== undefined) {
          return deviceData;
        }
      } else {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }

    return;
  };

  const validateInput = (): string | undefined => {
    if (name.length === 0) {
      return 'Please enter a name that is least 1 character long.';
    }
    if (!name.match(ValidDeviceNameChars)) {
      return 'Please only use letters and numbers in your device name.';
    }

    return undefined;
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    const validationError = validateInput();
    if (validationError !== undefined) {
      setErrorMessage(validationError);
    } else {
      try {
        setIsLoading(true);
        await addDevice();
        navigate(`${DEVICES_URL}`);
      } catch (e) {
        // Only set loading to false if the login failed. If we try to set it on a success, that gets
        // called after navigate which leads to the 'memory leak' React error.
        setIsLoading(false);
      }
    }
  };

  return (
    <Box sx={{
      textAlign: 'center',
      gridTemplateAreas: `
        "title"
        "error"
        "form"
      `,
    }}>
      <Typography
        variant="h3"
        sx={{
          padding: '12px',
          gridArea: 'title',
        }}>
        Add a Device
      </Typography>
      <Alert
        sx={{
          visibility: !!errorMessage ? 'visible' : 'hidden',
          gridArea: 'error',
        }}
        severity="error">
        {errorMessage}
      </Alert>
      <StyledForm
        sx={{
          gridArea: 'form',
          paddingTop: '18px',
          display: 'grid',
          rowGap: '12px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
            "name"
            "submit"
          `
        }}
        onSubmit={onSubmit} >
        <TextField
          sx={{
            gridArea: 'name'
          }}
          onChange={(e) => setName(e.target.value)}
          value={name}
          label="Device Name"
          id="devicename"
          type="text" />
        <LoadingButton loading={isLoading} sx={{ gridArea: 'submit' }} type="submit">
          <Typography variant="button">Create</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
