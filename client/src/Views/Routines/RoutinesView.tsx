import { Box, List, ListItem, ListItemText, Typography } from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import { GetRoutinesFetchURL, ParseRoutineArray, StoredRoutine } from '../../Utils/BackendIntegration';
import {
  useAuth,
} from '../../Utils/LoginState';

/**
 * The view used to describe the available routines for the user.
 * 
 * @returns The view.
 */
export default function RoutinesView() {
  const authState = useAuth();
  const [routines, setRoutines] = useState<StoredRoutine[]>([]);

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

  const routinesListItems: JSX.Element[] = useMemo(() => routines.map((r) => (
    <ListItem>
      <ListItemText primary={r.Name} />
    </ListItem>
  )), []);

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
      {/* <Box sx={{
        gridArea: 'routines',
      }}>
        {routineCards}
      </Box> */}
      <List sx={{ gridArea: 'routines' }}>
        {routinesListItems}
      </List>
    </Box>
  );
}
