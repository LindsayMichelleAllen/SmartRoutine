import {
  LoadingButton,
} from '@mui/lab';
import {
  Alert,
  Box,
  styled,
  Typography,
} from '@mui/material';
import React, {
  useEffect,
  useState,
} from 'react';
import {
  useSearchParams,
} from 'react-router-dom';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';
import {
  FetchRequest,
  ParseDevice,
} from '../../Utils/BackendIntegration';
import {
  DEVICE_ID_SEARCH_PARAM,
} from '../../Utils/CommonRouting';
import {
  ValidateDeviceName,
} from '../../Utils/InputValidation';

/**
 * A view used to edit a particular device.
 * 
 * @returns The view.
 */
export default function EditDeviceView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();

  const [name, setName] = useState('');

  const [nameValidation, setNameValidation] = useState('');

  const [isLoading, setIsLoading] = useState(false);
  const [genericError, setGenericError] = useState('');
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
      setGenericError(`${e}`);
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
      setGenericError(`${e}`);
    }
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    try {
      setIsLoading(true);
      await editDevice();
      await loadDevice();
      setSuccessMessage('Successfully updated the device!');
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
    } finally {
      setIsLoading(false);
    }
  };

  const onDeviceNameChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setName(input);
    const validationError = ValidateDeviceName(input);
    if (!!validationError) {
      setNameValidation(validationError);
    } else {
      setNameValidation('');
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
        visibility: !!genericError ? 'visible' : 'hidden',
      }}>
        {genericError}
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
        <ValidatedInput
          sx={{
            gridArea: 'name',
          }}
          labelId="name"
          labelText="Name"
          value={name}
          errorMessage={nameValidation}
          onValueChange={onDeviceNameChange} />
        <LoadingButton sx={{ gridArea: 'submit' }} type="submit" loading={isLoading}>
          <Typography variant="button">Update</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
