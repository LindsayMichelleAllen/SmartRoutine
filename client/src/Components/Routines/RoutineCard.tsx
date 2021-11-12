import { Avatar, Box, Card, CardContent, List, ListItem, ListItemAvatar, ListItemText, Typography } from '@mui/material';
import React, { useMemo } from 'react';
import { StoredRoutine } from '../../Utils/BackendIntegration';
import SettingsIcon from '@mui/icons-material/Settings';

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

  const listItems: JSX.Element[] = useMemo(() => routine.Configuration.map((c) => (
    <ListItem sx={{
    }}
      key={c.Id}>
      <ListItemAvatar>
        <Avatar>
          <SettingsIcon />
        </Avatar>
      </ListItemAvatar>
      <ListItemText primary={c.Id} secondary={`Offset: ${c.Offset}`} />
    </ListItem>
  )), [routine.Configuration]);

  return (
    <Card sx={{
      minWidth: '480px',
      textAlign: 'left',
    }}
      variant="outlined">
      <CardContent sx={{
        display: 'grid',
        gridTemplateAreas: `
          "routinetitle"
          "."
          "configurationstitle"
          "configurationlist"
        `,
      }}>
        <Typography sx={{ paddingBottom: '12px', gridArea: 'routinetitle' }} variant="h4">
          {routine.Name}
        </Typography>
        <Typography sx={{ gridArea: 'configurationstitle' }} variant="h6" color="text.secondary">
          Configurations
        </Typography>
        <List
          sx={{
            gridArea: 'configurationlist',
            display: 'flex',
            bgcolor: 'background.paper',
          }}>
          {listItems}
        </List>
      </CardContent>
    </Card>
  );
}
