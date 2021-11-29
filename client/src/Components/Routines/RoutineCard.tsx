import {
  Card,
  CardActions,
  CardContent,
  IconButton,
  Typography,
} from '@mui/material';
import React, { useMemo } from 'react';
import { StoredRoutine } from '../../Utils/BackendIntegration';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import VisibilityIcon from '@mui/icons-material/Visibility';
import { GetAlarmText } from '../../Views/Routines/RoutineUtils';

/**
 * The props for the {@link RoutineCard} component.
 */
export type RoutineCardProps = {
  /**
   * The routine that this card represents.
   */
  routine: StoredRoutine;

  /**
   * An event callback when a routine edit button is pressed.
   */
  onEditRoutine?: (routine: StoredRoutine) => void;

  /**
   * An event callback when a routine delete button is pressed.
   */
  onDeleteRoutine?: (routine: StoredRoutine) => void;

  /**
   * An event callback when a view button is pressed.
   */
  onViewRoutine?: (routine: StoredRoutine) => void;
}

/**
 * RoutineCard is a component used to display the information for a routine as well as some basic
 * actions available to the routine.
 * 
 * @param props See {@link RoutineCardProps}.
 * @returns The component.
 */
export default function RoutineCard(props: RoutineCardProps) {
  const {
    routine,
    onDeleteRoutine,
    onEditRoutine,
    onViewRoutine,
  } = props;

  const alarmText = useMemo(() =>
    GetAlarmText(routine?.BaseAlarm ?? new Date(0)
  ), [routine?.BaseAlarm]);

  const deviceCount = useMemo(() => {
    return routine?.Configuration.flatMap((c) => c.Device).filter((d) => !!d.Id).length;
  }, [routine?.Configuration]);

  const editButton = useMemo(() => !!onEditRoutine ? (
    <IconButton title="Edit" onClick={() => onEditRoutine(routine)}>
      <EditIcon />
    </IconButton>
  ) : (<></>), [onEditRoutine]);

  const deleteButton = useMemo(() => !!onDeleteRoutine ? (
    <IconButton title="Delete" onClick={() => onDeleteRoutine(routine)}>
      <DeleteIcon />
    </IconButton>
  ) : (<></>), [onDeleteRoutine]);

  const viewButton = useMemo(() => !!onViewRoutine ? (
    <IconButton title="View" onClick={() => onViewRoutine(routine)}>
      <VisibilityIcon />
    </IconButton>
  ) : (<></>), [onViewRoutine]);

  return (
    <Card>
      <CardContent sx={{
        display: 'grid',
        gridTemplateAreas: `
          "title title"
          "basealarm timeicon"
          "devicecount devicecount"
        `,
        justifyContent: 'center',
        alignItems: 'center',
        rowGap: '12px',
      }}>
        <Typography variant="h6" sx={{ gridArea: 'title' }}>
          {routine?.Name}
        </Typography>
        <Typography variant="body2" sx={{ gridArea: 'basealarm' }}>
          {alarmText}
        </Typography>
        <AccessTimeIcon sx={{ gridArea: 'timeicon' }} />
        <Typography variant="caption" sx={{ gridArea: 'devicecount' }}>
          {deviceCount} devices connected
        </Typography>
      </CardContent>
      <CardActions>
        {editButton}
        {deleteButton}
        {viewButton}
      </CardActions>
    </Card>
  );
}
