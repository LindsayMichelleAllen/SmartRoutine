import {
  LoadingButton,
} from '@mui/lab';
import {
  Box,
  styled,
  Typography,
} from '@mui/material';
import React, {
  useState,
} from 'react';
import {
  useNavigate,
} from 'react-router';
import AlertsBox from '../../Components/Containers/AlertsBox';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';
import {
  FetchRequest,
  ParseDevice,
} from '../../Utils/BackendIntegration';
import {
  DEVICES_URL,
} from '../../Utils/CommonRouting';
import {
  ValidateDeviceName,
} from '../../Utils/InputValidation';
import {
  useAuth,
} from '../../Utils/LoginState';

/**
 * AddDeviceView provides the means for a user to create a new device.
 * 
 * @returns The view.
 */
export default function AddDeviceView() {
  const [name, setName] = useState('');
  
  const [nameValidation, setNameValidation] = useState('');

  const [isLoading, setIsLoading] = useState(false);
  const [genericError, setGenericError] = useState('');

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
      setGenericError(`${e}`);
    }

    return;
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    try {
      setIsLoading(true);
      await addDevice();
      navigate(`${DEVICES_URL}`);
    } catch (e) {
      // Only set loading to false if the login failed. If we try to set it on a success, that gets
      // called after navigate which leads to the 'memory leak' React error.
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
      <AlertsBox errorMessage={genericError} />
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
        <ValidatedInput
          sx={{
            gridArea: 'name',
          }}
          labelId="device-id"
          labelText="Device Name"
          value={name}
          errorMessage={nameValidation}
          onValueChange={onDeviceNameChange} />
        <LoadingButton loading={isLoading} sx={{ gridArea: 'submit' }} type="submit">
          <Typography variant="button">Create</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
