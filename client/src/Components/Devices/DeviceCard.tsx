import { Card, CardActions, CardContent, IconButton, Typography } from '@mui/material';
import React from 'react';
import { StoredDevice } from '../../Utils/BackendIntegration';
import CloseIcon from '@mui/icons-material/Close';
import EditIcon from '@mui/icons-material/Edit';

export type DeviceCardProps = {
  device: StoredDevice;
  onDeleteDevice: (device: StoredDevice) => void;
  onEditDevice: (device: StoredDevice) => void;
}

/**
 * @param props
 */
export default function DeviceCard(props: DeviceCardProps) {
  const {
    device,
    onDeleteDevice,
    onEditDevice,
  } = props;

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
      <CardActions sx={{ justifyContent: 'end' }}>
        <IconButton title="Edit" onClick={() => onEditDevice(device)}>
          <EditIcon />
        </IconButton>
        <IconButton title="Delete" onClick={() => onDeleteDevice(device)}>
          <CloseIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
}
