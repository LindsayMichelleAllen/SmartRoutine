import { LoadingButton } from '@mui/lab';
import {
  Alert,
  Box,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import React, {
  useEffect,
  useState,
} from 'react';
import { useSearchParams } from 'react-router-dom';
import {
  FetchRequest,
  ParseDevice,
} from '../../Utils/BackendIntegration';
import { DEVICE_ID_SEARCH_PARAM } from '../../Utils/CommonRouting';
import { ValidDeviceNameChars } from '../../Utils/InputValidation';

/**
 * A view used to edit a particular device.
 * 
 * @returns The view.
 */
export default function EditDeviceView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();

  const [name, setName] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const deviceId = searchParams.get(DEVICE_ID_SEARCH_PARAM) ?? '';

  const loadDevice = async () => {
    try {
      const response = await FetchRequest('deviceRead', {
        id: deviceId,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const deviceData = ParseDevice(text);
      if (deviceData) {
        setName(deviceData.Name);
      } else {
        throw 'There was an error editing the device. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }
  };

  const editDevice = async () => {
    try {
      const response = await FetchRequest('deviceUpdate', {
        deviceid: deviceId,
        name,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }
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
    setIsLoading(true);

    const validationError = validateInput();
    if (validationError !== undefined) {
      setErrorMessage(validationError);
    } else {
      try {
        await editDevice();
        await loadDevice();
        setSuccessMessage('Successfully updated the device!');
      } catch (e) {
        console.error(e);
        setErrorMessage(`${e}`);
      } finally {
        setIsLoading(false);
      }
    }
  };

  useEffect(() => {
    loadDevice();
  }, [searchParams]);

  return (
    <Box sx={{
      textAlign: 'center',
      gridTemplateAreas: `
        "title"
        "error"
        "success"
        "form"
      `,
    }}>
      <Typography sx={{ gridArea: 'title' }} variant="h3">
        Update a device
      </Typography>
      <Alert severity="error" sx={{
        gridArea: 'error',
        visibility: !!errorMessage ? 'visible' : 'hidden',
      }}>
        {errorMessage}
      </Alert>
      <Alert severity="success" sx={{
        gridArea: 'success',
        visibility: !!successMessage ? 'visible' : 'hidden',
      }}>
        {successMessage}
      </Alert>
      <StyledForm
        onSubmit={onSubmit}
        sx={{
          gridArea: 'form',
          display: 'grid',
          rowGap: '12px',
          paddingTop: '18px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
          "name"
          "submit"
        `,
        }}>
        <TextField sx={{
          gridArea: 'name'
        }}
          value={name}
          onChange={(e) => setName(e.target.value)}
          label="Device Name"
          id="devicename"
          type="text" />
        <LoadingButton sx={{ gridArea: 'submit' }} type="submit" loading={isLoading}>
          <Typography variant="button">Update</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
