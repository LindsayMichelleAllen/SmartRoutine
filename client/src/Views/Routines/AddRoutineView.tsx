import {
  Box,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import {
  LoadingButton,
  TimePicker,
} from '@mui/lab';
import React, {
  useState,
} from 'react';
import {
  FetchRequest,
  GetRoutineBasealarmString,
  ParseRoutine,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import {
  useAuth,
} from '../../Utils/LoginState';
import {
  useNavigate,
} from 'react-router';
import {
  ValidateRoutineName,
} from '../../Utils/InputValidation';
import {
  ROUTINES_URL,
} from '../../Utils/CommonRouting';
import AlertsBox from '../../Components/Containers/AlertsBox';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';

/**
 * A view used for the user to create a new routine.
 * 
 * @returns The view.
 */
export default function AddRoutineView() {
  const [name, setName] = useState('');
  const [time, setTime] = useState<Date>(new Date(Date.now()));

  const [nameValidation, setNameValidation] = useState('');

  const [isLoading, setIsLoading] = useState(false);
  const [genericError, setGenericError] = useState('');
  
  const navigate = useNavigate();

  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const addRoutine = async (): Promise<StoredRoutine | undefined> => {
    try {
      const timeString = GetRoutineBasealarmString(time);
      const response = await FetchRequest('routineCreate', {
        name,
        userid: loginDetails.Username,
        basealarm: timeString,
      });

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
      setGenericError(`${e}`);
      throw e;
    }
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    try {
      setIsLoading(true);
      await addRoutine();
      navigate(`${ROUTINES_URL}`);
    } catch (e) {
      // Only set loading to false if the login failed. If we try to set it on a success, that gets
      // called after navigate which leads to the 'memory leak' React error.
      setIsLoading(false);
    }
  };

  const onRoutineNameChange: OnValidatedInputChange = (e) => {
    const input = e.target.value;
    setName(input);
    const validationError = ValidateRoutineName(input);
    if (!!validationError) {
      setNameValidation(validationError);
    } else {
      setNameValidation('');
    }
  };

  return (
    <Box
      sx={{
        display: 'grid',
        gridTemplateAreas: `
          "title"
          "alerts-box"
          "routines"
        `,
        textAlign: 'center',
      }}>
      <Typography sx={{ gridArea: 'title' }} variant="h3" >
        Add a Routine
      </Typography>
      <AlertsBox errorMessage={genericError} />
      <StyledForm
        onSubmit={onSubmit}
        sx={{
          paddingTop: '18px',
          display: 'grid',
          rowGap: '12px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
            "name"
            "time"
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
          onValueChange={onRoutineNameChange} />
        <TimePicker
          label="Base Alarm"
          value={time}
          onChange={(v) => setTime(v)}
          renderInput={(params) => <TextField sx={{ gridArea: 'time' }} {...params} />} />
        <LoadingButton loading={isLoading} sx={{ gridArea: 'submit' }} type="submit">
          <Typography variant="button">Create</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
