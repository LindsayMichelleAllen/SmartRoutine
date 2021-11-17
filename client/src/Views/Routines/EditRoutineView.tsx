import { TimePicker } from '@mui/lab';
import {
  Box,
  Typography,
  Alert,
  TextField,
  Button,
  CircularProgress,
  styled,
  List,
  ListItem,
} from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import {
  GetFetchRequest,
  GetGetRoutineURL,
  GetRoutineBasealarmString,
  GetUpdateRoutineURL,
  ParseRoutine,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import { ROUTINE_ID_SEARCH_PARAM } from '../../Utils/CommonRouting';
import { ValidRoutineNameChars } from '../../Utils/InputValidation';

/**
 * A view used for a user to edit the details for a given view.
 * 
 * @returns The View.
 */
export default function EditRoutineView() {
  const [name, setName] = useState('');
  const [time, setTime] = useState<Date>(new Date(Date.now()));
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);
  
  const devices = useMemo(() => routine?.Configuration?.map((c) => (
    <ListItem key={c.Id}>
      {c.Device.Name}
    </ListItem>
  )), [routine?.Configuration]);

  const routineId = searchParams.get(ROUTINE_ID_SEARCH_PARAM) ?? '';

  const loadRoutine = async () => {
    try {
      const response = await fetch(
        GetGetRoutineURL(),
        GetFetchRequest({
          routineid: routineId,
        }),
      );

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const routineData = ParseRoutine(text);
      if (routineData) {
        setRoutine(routineData);
        setName(routineData.Name);
        setTime(routineData.BaseAlarm);
      } else {
        throw 'There was an error editing the routine. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(e);
    }
  };

  const editRoutine = async () => {
    try {
      const timeString = GetRoutineBasealarmString(time);
      const response = await fetch(
        GetUpdateRoutineURL(),
        GetFetchRequest({
          name,
          routineid: routine.Id,
          basealarm: timeString,
        }),
      );

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(e);
      throw e;
    }
  };

  const validateInput = (): string | undefined => {
    if (!name.match(ValidRoutineNameChars)) {
      return 'Please only use letters and numbers in your routine name.';
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
        await editRoutine();
        await loadRoutine();
        setSuccessMessage('Successfully updated the routine!');
      } catch (e) {
        console.error(e);
      }
      finally {
        setIsLoading(false);
      }
    }
  };

  useEffect(() => {
    loadRoutine();
  }, [searchParams]);

  return (
    <Box>
      <Typography padding="12px" textAlign="center" variant="h3">
        Update a Routine
      </Typography>
      <StyledForm
        onSubmit={onSubmit}
        sx={{
          paddingTop: '18px',
          display: 'grid',
          rowGap: '12px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
            "error"
            "success"
            "name"
            "time"
            "devicestitle"
            "devices"
            "submit"
          `,
        }}>
        <Alert sx={{
          visibility: !!errorMessage ? 'visible' : 'hidden',
          gridArea: 'error',
        }} severity="error">
          {errorMessage}
        </Alert>
        <Alert sx={{
          visibility: !!successMessage ? 'visible' : 'hidden',
          gridArea: 'success',
        }} severity="success">
          {successMessage}
        </Alert>
        <TextField
          sx={{ gridArea: 'name' }}
          onChange={(e) => setName(e.target.value)}
          value={name}
          label="Routine Name"
          id="routinename"
          type="text" />
        <TimePicker
          label="Base Alarm"
          value={time}
          onChange={(v) => setTime(v)}
          renderInput={(params) => <TextField sx={{ gridArea: 'time' }} {...params} />} />
        <Typography variant="h6" sx={{ gridArea: 'devicestitle' }}>
          Devices
        </Typography>
        <List sx={{ gridArea: 'devices' }}>
          {devices}
        </List>
        <Button sx={{ gridArea: 'submit' }} type="submit" >
          {
            isLoading
              ? (<CircularProgress />)
              : (<Typography variant="button">Update</Typography>)
          }
        </Button>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
