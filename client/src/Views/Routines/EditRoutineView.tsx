import {
  LoadingButton,
  TimePicker,
} from '@mui/lab';
import {
  Box,
  Typography,
  TextField,
  styled,
} from '@mui/material';
import React, {
  useEffect,
  useState,
} from 'react';
import {
  useSearchParams,
} from 'react-router-dom';
import AlertsBox from '../../Components/Containers/AlertsBox';
import ValidatedInput, {
  OnValidatedInputChange,
} from '../../Components/Containers/ValidatedInput';
import {
  FetchRequest,
  GetRoutineBasealarmString,
  ParseRoutine,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import {
  ROUTINE_ID_SEARCH_PARAM,
} from '../../Utils/CommonRouting';
import {
  ValidateRoutineName,
} from '../../Utils/InputValidation';

/**
 * A view used for a user to edit the details for a given view.
 * 
 * @returns The View.
 */
export default function EditRoutineView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();

  const [name, setName] = useState('');
  const [time, setTime] = useState<Date>(new Date(Date.now()));

  const [nameValidation, setNameValidation] = useState('');

  const [isLoading, setIsLoading] = useState(false);
  const [genericError, setGenericError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);

  const routineId = searchParams.get(ROUTINE_ID_SEARCH_PARAM) ?? '';

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
        setName(routineData.Name);
        setTime(routineData.BaseAlarm);
      } else {
        throw 'There was an error editing the routine. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
    }
  };

  const editRoutine = async () => {
    try {
      const timeString = GetRoutineBasealarmString(time);
      const response = await FetchRequest('routineUpdate', {
        name,
        routineid: routine.Id,
        basealarm: timeString,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
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
      await editRoutine();
      await loadRoutine();
      setSuccessMessage('Successfully updated the routine!');
    } catch (e) {
      console.error(e);
    }
    finally {
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

  useEffect(() => {
    loadRoutine();
  }, [searchParams]);

  return (
    <Box>
      <Typography padding="12px" textAlign="center" variant="h3">
        Update a Routine
      </Typography>
      <AlertsBox errorMessage={genericError} successMessage={successMessage} />
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
        <LoadingButton sx={{ gridArea: 'submit' }} type="submit" loading={isLoading}>
          <Typography variant="button">Update</Typography>
        </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
