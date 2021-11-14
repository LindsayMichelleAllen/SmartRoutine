import {
  Alert,
  Box,
  Button,
  CircularProgress,
  styled,
  TextField,
  Typography,
} from '@mui/material';
import React, { useState } from 'react';
import {
  GetCreateRoutineURL,
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

/**
 * A view used for the user to create a new routine.
 * 
 * @returns The view.
 */
export default function AddRoutineView() {
  const [name, setName] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();
  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const addRoutine = async (): Promise<StoredRoutine | undefined> => {
    try {
      const response = await fetch(GetCreateRoutineURL(), {
        method: 'POST',
        body: `name=${name}&userid=${loginDetails?.userid ?? ''}`,
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
    setIsLoading(true);

    const validationError = validateInput();
    if (validationError !== undefined) {
      setErrorMessage(validationError);
    } else {
      try {
        const routine = await addRoutine();
        navigate(`/routine?routineid=${routine.Id}`);
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
              : (<Typography variant="button">Create</Typography>)
          }
        </Button>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
