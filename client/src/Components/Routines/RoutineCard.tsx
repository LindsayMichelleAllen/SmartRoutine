import {
  Card,
  CardActions,
  CardContent,
  IconButton,
  Typography,
} from '@mui/material';
import React from 'react';
import { StoredRoutine } from '../../Utils/BackendIntegration';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { useNavigate } from 'react-router-dom';
import { EDIT_ROUTINE_URL } from '../../Utils/CommonRouting';

/**
 * The props for the {@link RoutineCard} component.
 */
export type RoutineCardProps = {
  /**
   * The routine that this card represents.
   */
  routine: StoredRoutine;
}

/**
 * RoutineCard is a cpomoment used to display the information for a routine as well as some basic
 * actions available to the routine.
 * 
 * @param props See {@link RoutineCardProps}.
 * @returns The component.
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
