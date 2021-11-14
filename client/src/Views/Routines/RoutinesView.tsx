import {
  Box,
  Fab,
  Typography,
} from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import RoutineCard from '../../Components/Routines/RoutineCard';
import {
  GetRoutinesFetchURL,
  ParseRoutineArray,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import AddIcon from '@mui/icons-material/Add';
import {
  useAuth,
} from '../../Utils/LoginState';
import { useNavigate } from 'react-router';
import { ADD_ROUTINE_URL } from '../../Utils/CommonRouting';

/**
 * The view used to describe the available routines for the user.
 * 
 * @returns The view.
 */
export default function RoutinesView() {
  const [routines, setRoutines] = useState<StoredRoutine[]>([]);

  const navigate = useNavigate();
  const authState = useAuth();

  const loginDetails = authState?.loginDetails;

  const fetchRoutines = async () => {
    try {
      const response = await fetch(GetRoutinesFetchURL(), {
        method: 'POST',
        body: `userid=${loginDetails.userid ?? ''}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      });

      const text = await response.text();

      if (response.ok) {
        const routinesData = await ParseRoutineArray(text);
        if (routinesData !== undefined) {
          setRoutines(routinesData);
        }
      }
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    if (loginDetails !== undefined) {
      fetchRoutines();
    }
  }, [loginDetails]);

  console.log(routines);

  const routineCards = useMemo(() => routines.map((r) => (
    <RoutineCard routine={r} />
  )), [routines]);

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "title"
        "routines"
      `,
      gridTemplateRows: 'min-content 1fr',
      rowGap: '48px',
      textAlign: 'center',
    }}>
      <Typography
        sx={{ gridArea: 'title' }}
        variant="h2">
        Routines
      </Typography>
      <Box sx={{
        gridArea: 'routines',
        display: 'grid',
        columnGap: '12px',
        rowGap: '12px',
        padding: '12px',
        justifyContent: 'center',
        paddingBottom: '128px', // Add some extra space so the FAB doesn't overlay the actions.
        gridTemplateColumns: {
          sm: '220px 220px',
          xs: '1fr',
        },
        width: {
          sm: 'auto',
        }
      }}>
        {routineCards}
      </Box>
      <Fab sx={{
        position: 'absolute',
        bottom: '24px',
        right: '24px',
      }}
        color='primary'
        onClick={() => navigate(ADD_ROUTINE_URL)}>
        <AddIcon />
      </Fab>
    </Box>
  );
}
