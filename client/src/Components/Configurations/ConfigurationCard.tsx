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
  StoredConfiguration,
} from '../../Utils/BackendIntegration';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

/**
 * The props for {@link ConfigurationCard}.
 */
export type ConfigurationCardProps = {
  /**
   * The configuration that is represented by this card.
   */
  configuration: StoredConfiguration;

  /**
   * The callback event when the delete button is pressed for this configuration.
   */
  onDeleteConfiguration?: (configuration: StoredConfiguration) => void;

  /**
   * The callback event when the edit button is pressed for this configuration.
   */
  onEditConfiguration?: (configuration: StoredConfiguration) => void;
}

/**
 * ConfigurationCard is a component used to display the information for a configuration as well as
 * some basic actions available to the configuration.
 * 
 * @param props See {@link ConfigurationCardProps}.
 * @returns The component.
 */
export default function ConfigurationCard(props: ConfigurationCardProps) {
  const {
    configuration,
    onDeleteConfiguration,
    onEditConfiguration,
  } = props;

  const editButton = useMemo(() => !!onEditConfiguration ? (
    <IconButton title="Edit" onClick={() => onEditConfiguration(configuration)}>
      <EditIcon />
    </IconButton>
  ) : (<></>), [onEditConfiguration]);

  const deleteButton = useMemo(() => !!onDeleteConfiguration ? (
    <IconButton title="Delete" onClick={() => onDeleteConfiguration(configuration)}>
      <DeleteIcon />
    </IconButton>
  ) : (<></>), [onDeleteConfiguration]);

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
      <CardActions>
        {editButton}
        {deleteButton}
      </CardActions>
    </Card>
  );
}
