import { Typography } from '@mui/material';
import { Box } from '@mui/system';
import React from 'react';
import { useSearchParams } from 'react-router-dom';
import { RoutineIdSearchParam } from '../../Utils/CommonRouting';

/**
 * A view used to describe a single routine to the user.
 * 
 * @returns The view.
 */
export default function SingleRoutineView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  console.log(searchParams);

  const routineId = searchParams.get(RoutineIdSearchParam) ?? '';

  return (
    <Box>
      <Typography>
        {routineId}
      </Typography>
    </Box>
  );
}
