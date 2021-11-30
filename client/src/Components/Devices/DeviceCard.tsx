import {
  Card,
  CardActions,
  CardContent,
  IconButton,
  Typography,
} from '@mui/material';
import React, {
  useMemo,
} from 'react';
import {
  StoredDevice,
} from '../../Utils/BackendIntegration';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

/**
 * The props for {@link DeviceCard}.
 */
export type DeviceCardProps = {
  /**
   * The device to be represented by this card.
   */
  device: StoredDevice;

  /**
   * The callback event for when the delete button is pressed on this card.
   */
  onDeleteDevice: (device: StoredDevice) => void;

  /**
   * The callback event for when the edit button is pressed on this card.
   */
  onEditDevice: (device: StoredDevice) => void;
}

/**
 * A card component used to represent a single device.
 * 
 * @param props See {@link DeviceCardProps}.
 * @returns The component.
 */
export default function DeviceCard(props: DeviceCardProps) {
  const {
    device,
    onDeleteDevice,
    onEditDevice,
  } = props;

  const editButton = useMemo(() => !!onEditDevice ? (
    <IconButton title="Edit" onClick={() => onEditDevice(device)}>
      <EditIcon />
    </IconButton>
  ) : (<></>), [onEditDevice]);

  const deleteButton = useMemo(() => !!onDeleteDevice ? (
    <IconButton title="Delete" onClick={() => onDeleteDevice(device)}>
      <DeleteIcon />
    </IconButton>
  ) : (<></>), [onDeleteDevice]);

  return (
    <Card>
      <CardContent sx={{
        display: 'grid',
        gridTemplateAreas: `
          "deviceName"
        `
      }}>
        <Typography variant="h6" sx={{ gridArea: 'deviceName' }}>
          {device?.Name ?? ''}
        </Typography>
      </CardContent>
      <CardActions>
        {editButton}
        {deleteButton}
      </CardActions>
    </Card>
  );
}
