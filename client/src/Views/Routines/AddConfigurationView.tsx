import {
  LoadingButton,
} from '@mui/lab';
import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import {
  Box,
} from '@mui/system';
import React, {
  useEffect,
  useMemo,
  useState,
} from 'react';
import {
  useNavigate,
  useSearchParams,
} from 'react-router-dom';
import AlertsBox from '../../Components/Containers/AlertsBox';
import {
  FetchRequest,
  ParseDeviceArray,
  ParseRoutine,
  StoredDevice,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import {
  ROUTINE_ID_SEARCH_PARAM,
  VIEW_ROUTINE_URL,
} from '../../Utils/CommonRouting';
import {
  useAuth,
} from '../../Utils/LoginState';

/**
 * A view used to add a new device configuration to an existing routine.
 * 
 * @returns The view.
 */
export default function AddConfigurationView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [deviceId, setDeviceId] = useState('');
  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);
  const [deviceList, setDeviceList] = useState<StoredDevice[]>([]);
  const [genericError, setGenericError] = useState('');
  const [offset, setOffset] = useState(0);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const navigate = useNavigate();
  const authState = useAuth();
  const loginDetails = authState?.loginDetails;
  const routineId = searchParams.get(ROUTINE_ID_SEARCH_PARAM) ?? '';

  const loadDeviceOptions = async () => {
    try {
      const response = await FetchRequest('deviceReadUser', {
        userid: loginDetails.Username,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const devices = ParseDeviceArray(text);
      if (devices) {
        setDeviceList(devices);
      } else {
        throw 'There was an error fetching the device list. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
    }
  };

  const loadRoutine = async () => {
    try {
      const response = await FetchRequest('routineRead', {
        routineid: routineId,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const routineData = ParseRoutine(text);
      if (routineData) {
        setRoutine(routineData);
      } else {
        throw 'There was an error editing the routine. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
    }
  };

  const addConfiguration = async () => {
    try {
      const response = await FetchRequest('configurationCreate', {
        deviceid: deviceId,
        offset: `${offset}`,
        routineid: routineId,
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

  const validateInput = (): string | undefined => {
    if (typeof offset !== 'number') {
      return 'Please enter a number for the offset value.';
    }
    if (!deviceId) {
      return 'Please select a device before submitting.';
    }

    return undefined;
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    const validationError = validateInput();
    if (!!validationError) {
      setGenericError(validationError);
      setIsSubmitting(false);
    } else {
      try {
        await addConfiguration();
        navigate(`${VIEW_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routineId}`);
      } catch (e) {
        console.error(e);
        setGenericError(`${e}`);
        setIsSubmitting(false);
      }
    }
  };

  const configurationOptions = useMemo(() => deviceList.map((d) => (
    <MenuItem key={d.Id} value={d.Id}>
      {d.Name}
    </MenuItem>
  )), [deviceList]);

  useEffect(() => {
    loadRoutine().then(() => {
      loadDeviceOptions();
    });
  }, []);

  const handleSelect = (event: SelectChangeEvent<string>) => {
    const selectedDevice = deviceList.find((d) => d.Id === event.target.value);
    setDeviceId(selectedDevice.Id);
  };

  return (
    <Box sx={{
      textAlign: 'center',
      display: 'grid',
      gridTemplateAreas: `
        "title"
        "alerts"
        "form"
      `
    }}>
      <Typography
        sx={{
          gridArea: 'title',
        }}
        variant="h3">
        Add a Device to {routine?.Name ?? ''}
      </Typography>
      <AlertsBox errorMessage={genericError} />
      <StyledForm
        onSubmit={onSubmit}
        sx={{
          gridArea: 'form',
          paddingTop: '18px',
          display: 'grid',
          rowGap: '12px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
            "device-select"
            "offset"
            "submit"
          `,
        }}>
        <FormControl sx={{ gridArea: 'device-select' }}>
          <InputLabel id="configuration-device-selection-label">Device</InputLabel>
          <Select
            labelId="configuration-device-selection-label"
            id="configuration-device-selection"
            value={deviceId}
            onChange={handleSelect}>
            {configurationOptions}
          </Select>
        </FormControl>
        <TextField
          value={offset}
          onChange={(e) => {
            const offsetInput = e.target.value;
            if (/^[0-9]+$/.test(offsetInput)) {
              const parsedInput = parseInt(offsetInput);
              const valueToSet = parsedInput === NaN ? 0 : parsedInput;
              setOffset(valueToSet);
            }
          }}
          label="Device Offset"
          id="device-offset"
          type="number"
          sx={{
            gridArea: 'offset'
          }} />
        <LoadingButton sx={{ gridArea: 'submit' }} type="submit" loading={isSubmitting}>
          <Typography variant="button">Add this device</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
