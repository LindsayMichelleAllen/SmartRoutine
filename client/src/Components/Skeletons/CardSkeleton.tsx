import { Skeleton } from '@mui/material';
import React from 'react';

/**
 * A component used to represent a loading card.
 * 
 * @returns The component.
 */
export default function CardSkeleton() {
  return (
    <Skeleton
      variant="rectangular"
      sx={{
        minWidth: {
          xs: '100%',
          sm: '320px',
        },
        minHeight: '200px',
      }}
    />
  );
}
