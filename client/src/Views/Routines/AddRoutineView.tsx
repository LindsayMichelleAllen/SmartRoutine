import {
  Alert,
  Box,
  Button,
  CircularProgress,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import {
  TimePicker,
} from '@mui/lab';
import React, { useState } from 'react';
import {
  GetCreateRoutineURL,
  GetFetchRequest,
  GetRoutineBasealarmString,
  ParseRoutine,
  StoredRoutine,
 } from '../../Utils/BackendIntegration';
import { useAuth } from '../../Utils/LoginState';
import {
  useNavigate,
} from 'react-router';
import {
  ValidRoutineNameChars,
} from '../../Utils/InputValidation';
import { EDIT_ROUTINE_URL, ROUTINES_URL, ROUTINE_ID_SEARCH_PARAM } from '../../Utils/CommonRouting';

/**
 * A view used for the user to create a new routine.
 * 
 * @returns The view.
 */
export default function AddRoutineView() {
  const [name, setName] = useState('');
  const [time, setTime] = useState<Date>(new Date(Date.now()));
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();
  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const addRoutine = async (): Promise<StoredRoutine | undefined> => {
    try {
      const timeString = GetRoutineBasealarmString(time);
      const response = await fetch(
        GetCreateRoutineURL(),
        GetFetchRequest({
          name,
          userid: loginDetails.Username,
          basealarm: timeString,
        }),
      );

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const routineData = ParseRoutine(text);
      if (routineData) {
        return routineData;
      } else {
        throw 'There was an error creating the routine. Please try again.';
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

    const validationError = validateInput();
    if (validationError !== undefined) {
      setErrorMessage(validationError);
    } else {
      try {
        setIsLoading(true);
        await addRoutine();
        navigate(`${ROUTINES_URL}`);
      } catch (e) {
        // Only set loading to false if the login failed. If we try to set it on a success, that gets
        // called after navigate which leads to the 'memory leak' React error.
        setIsLoading(false);
      }
    }
  };

  return (
    <Box>
      <Typography padding="12px" textAlign="center" variant="h3">
        Add a Routine
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
            "name"
            "time"
            "submit"
          `,
        }}>
        <Alert sx={{ visibility: !!errorMessage ? 'visible' : 'hidden' }} severity="error">
          {errorMessage}
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
        <Button sx={{ gridArea: 'submit' }} type="submit" >
          {
            isLoading
              ? (<CircularProgress />)
              : (<Typography variant="button">Create</Typography>)
          }
        </Button>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
