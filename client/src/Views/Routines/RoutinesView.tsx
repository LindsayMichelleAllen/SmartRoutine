import { Box, Typography } from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import RoutineCard from '../../Components/Routines/RoutineCard';
import { GetRoutinesFetchURL, ParseRoutineArray, StoredRoutine } from '../../Utils/BackendIntegration';
import { useLoginState } from '../../Utils/LoginState';

/**
 * The view used to describe the available routines for the user.
 * 
 * @returns The view.
 */
export default function RoutinesView() {
  const loginState = useLoginState();
  const [routines, setRoutines] = useState<StoredRoutine[]>([]);

  const getRoutines = async () => {
    try {
      await fetch(GetRoutinesFetchURL(), {
        method: 'POST',
        body: `userid=${loginState.userid ?? ''}`,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
      })
        .then((response) => {
          if (response.ok) {
            response.text().then((data) => {
              const routinesData = ParseRoutineArray(data);
              if (routinesData !== undefined) {
                setRoutines(routinesData);
              } else {
                console.error('There was an error parsing the stored routines.');
              }
            });
          } else {
            console.error('There was an error parsing the stored routines.');
          }
        });
    } catch (e) {
      console.error(e);
    }
  };

  useEffect(() => {
    if (loginState !== undefined) {
      getRoutines();
    }
  }, [loginState]);

  console.log(routines);

  const routineCards: JSX.Element[] = useMemo(() => routines.map((r) => (
    <RoutineCard routine={r} key={r.Id} />
  )), [routines]);

  return (
    <Box sx={{
      height: '100%',
      width: '100%',
      display: 'grid',
      gridTemplateAreas: `
        "title title title"
        ". routines ."
      `,
      gridTemplateRows: 'min-content 1fr',
      gridTemplateColumns: '1fr max-content 1fr',
      rowGap: '48px',
      textAlign: 'center',
    }}>
      <Typography
        sx={{
          gridArea: 'title'
        }}
        variant="h2">
        Routines
      </Typography>
      <Box sx={{
        gridArea: 'routines',
      }}>
        {routineCards}
      </Box>
    </Box>
  );
}
