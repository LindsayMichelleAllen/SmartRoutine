import { Typography } from '@mui/material';
import { Box } from '@mui/system';
import React from 'react';
import { useSearchParams } from 'react-router-dom';
import { RoutineIdSearchParam } from '../../Utils/CommonRouting';

export type SingleRoutineViewProps = {
}

/**
 * @param props
 */
export default function SingleRoutineView(props: SingleRoutineViewProps) {
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
