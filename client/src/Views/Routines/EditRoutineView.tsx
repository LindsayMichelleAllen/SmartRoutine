import {
  Box,
  Typography,
  Alert,
  TextField,
  Button,
  CircularProgress,
  styled,
} from '@mui/material';
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import { useSearchParams } from 'react-router-dom';
import {
  GetGetRoutineURL,
  GetUpdateRoutineURL,
  ParseRoutine,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import { RoutineIdSearchParam } from '../../Utils/CommonRouting';
import { ValidRoutineNameChars } from '../../Utils/InputValidation';

/**
 * A view used for a user to edit the details for a given view.
 * 
 * @returns The View.
 */
export default function EditRoutineView() {
  const [name, setName] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);
  const navigate = useNavigate();

  const routineId = searchParams.get(RoutineIdSearchParam) ?? '';

  const loadRoutine = async () => {
    try {
      const response = await fetch(GetGetRoutineURL(), {
        method: 'POST',
        body: `routineid=${routineId}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const routineData = ParseRoutine(text);
      if (routineData) {
        setRoutine(routineData);
        setName(routineData.Name);
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
      const response = await fetch(GetUpdateRoutineURL(), {
        method: 'POST',
        body: `routineId=${routine.Id}}&name=${name}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      });

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
        navigate(`/routine?routineid=${routine.Id}`);
      } catch (e) {
        // Only set loading to false if the login failed. If we try to set it on a success, that gets
        // called after navigate which leads to the 'memory leak' React error.
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
          gridTemplateAreas: `
            "error"
            "name"
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
