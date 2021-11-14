import { Avatar, Box, Button, Card, CardActions, CardContent, IconButton, List, ListItem, ListItemAvatar, ListItemText, Typography } from '@mui/material';
import React, { useMemo } from 'react';
import { GetDeleteRoutineURL, StoredRoutine } from '../../Utils/BackendIntegration';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { useNavigate } from 'react-router-dom';
import { EDIT_ROUTINE_URL } from '../../Utils/CommonRouting';

export type RoutineCardProps = {
  routine: StoredRoutine;
}

/**
 * @param props
 */
export default function RoutineCard(props: RoutineCardProps) {
  const {
    routine,
  } = props;

  const navigate = useNavigate();

  return (
    <Card sx={{
      maxWidth: {
        sm: '480px',
        xs: '100%',
      },
    }}
      variant="outlined"
    >
      <CardContent sx={{
      }}>
        <Typography variant="h6" sx={{ gridArea: 'routine_title' }}>
          {routine.Name}
        </Typography>
        <Typography variant="caption" sx={{ gridArea: 'device_count' }}>
          {routine.Configuration.length} devices connected
        </Typography>
      </CardContent>
      <CardActions sx={{
        justifyContent: 'end',
      }}>
        <IconButton title="Edit" onClick={() => navigate(`${EDIT_ROUTINE_URL}?routineid=${routine.Id}`)}>
          <EditIcon />
        </IconButton>
        <IconButton title="Delete">
          <DeleteIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}
