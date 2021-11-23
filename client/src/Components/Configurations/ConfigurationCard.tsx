import { Card, CardActions, CardContent, IconButton, Typography } from '@mui/material';
import React from 'react';
import { StoredConfiguration } from '../../Utils/BackendIntegration';
import CloseIcon from '@mui/icons-material/Close';
import EditIcon from '@mui/icons-material/Edit';

export type ConfigurationCardProps = {
  configuration: StoredConfiguration;
  onDeleteConfiguration: (configuration: StoredConfiguration) => void;
  onEditConfiguration: (configuration: StoredConfiguration) => void;
}

/**
 * @param props
 */
export default function ConfigurationCard(props: ConfigurationCardProps) {
  const {
    configuration,
    onDeleteConfiguration,
    onEditConfiguration,
  } = props;

  return (
    <Card>
      <CardContent sx={{
        display: 'grid',
        gridTemplateAreas: `
          "deviceName"
          "offset"
        `,
      }}>
        <Typography variant="h6" sx={{ gridArea: 'deviceName' }}>
          {configuration.Device?.Name ?? ''}
        </Typography>
        <Typography variant="body2" sx={{ gridArea: 'offset' }}>
          Offset: {configuration?.Offset ?? 0}
        </Typography>
      </CardContent>
      <CardActions sx={{ justifyContent: 'end' }}>
        <IconButton title="Edit" onClick={() => onEditConfiguration(configuration)}>
          <EditIcon />
        </IconButton>
      </CardActions>
      <IconButton title="Close" onClick={() => onDeleteConfiguration(configuration)}>
        <CloseIcon />
      </IconButton>
    </Card>
  );
}
