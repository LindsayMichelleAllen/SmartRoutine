import {
  Box,
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogTitle,
  Fab,
  IconButton,
  Paper,
  Typography,
} from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { FetchRequest, ParseRoutine, StoredConfiguration, StoredRoutine } from '../../Utils/BackendIntegration';
import { ADD_ROUTINE_URL, ROUTINE_ID_SEARCH_PARAM } from '../../Utils/CommonRouting';
import ConfigurationCard from '../../Components/Configurations/ConfigurationCard';
import EditIcon from '@mui/icons-material/Edit';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import AddIcon from '@mui/icons-material/Add';
import { GetAlarmText } from './RoutineUtils';

/**
 *
 */
export default function SingleRoutineView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);
  const [alarmText, setAlarmText] = useState<string>('');
  const [errorMessage, setErrorMessage] = useState('');
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleteProcessing, setDeleteProcessing] = useState(false);
  const [configToDelete, setConfigToDelete] = useState<StoredConfiguration | undefined>(undefined);

  const navigate = useNavigate();
  const routineId = searchParams.get(ROUTINE_ID_SEARCH_PARAM) ?? '';

  const loadRoutine = async () => {
    try {
      const response = await FetchRequest('routineRead', {
        routineid: routineId,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      } else {
        const routine = ParseRoutine(text);
        setRoutine(routine);
        setAlarmText(GetAlarmText(routine.BaseAlarm));
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }
  };

  const deleteConfig = async () => {
    try {
      const response = await FetchRequest('configurationDelete', {
        id: configToDelete.Id,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(`${e}`);
    }
  };

  useEffect(() => {
    loadRoutine();
  }, [routineId]);

  const handleDeleteConfiguration = (configuration: StoredConfiguration) => {
    setDeleteDialogOpen(true);
    setConfigToDelete(configuration);
  };

  const handleEditConfiguration = (configuration: StoredConfiguration) => {
    navigate(`${ADD_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routine.Id}`);
  };

  const onDeleteConfiguration = async () => {
    setDeleteProcessing(true);
    await deleteConfig();
    setDeleteProcessing(false);

    setDeleteDialogOpen(false);
    await loadRoutine();
  };

  const configurations = useMemo(() => routine?.Configuration?.map((c) => (
    <ConfigurationCard
      configuration={c}
      onDeleteConfiguration={handleDeleteConfiguration}
      onEditConfiguration={handleEditConfiguration}
    />
  )), [routine?.Configuration]);

  return (
    <Box sx={{ textAlign: 'center' }}>
      <Box sx={{
        display: 'grid',
        gridTemplateAreas: `
          "routineDetails"
          "configurationsTitle"
          "configurationsList"
        `,
      }}>
        <Paper
          elevation={2}
          sx={{
            position: 'relative',
            gridArea: 'routineDetails',
            display: 'grid',
            margin: '18px',
            padding: '12px',
            alignItems: 'center',
            justifyContent: 'center',
            gridTemplateAreas: `
            "title title"
            "alarm alarmIcon"
          `
          }}>
          <Typography sx={{ gridArea: 'title' }} padding="12px" textAlign="center" variant="h3">
            {routine?.Name ?? ''}
          </Typography>
          <Typography sx={{ gridArea: 'alarm' }}>
            {alarmText}
          </Typography>
          <AccessTimeIcon sx={{ gridArea: 'alarmIcon' }} />
          <IconButton sx={{ position: 'absolute', right: '18px', bottom: '0px' }} title="Edit">
            <EditIcon />
          </IconButton>
        </Paper>
      </Box>
      <Typography
        sx={{
          gridArea: 'configurationsTitle',
        }}
        variant="h4">
        Devices
      </Typography>
      <Box sx={{
        gridArea: 'configurationsList',
        display: 'grid',
        columnGap: '12px',
        rowGap: '12px',
        padding: '12px',
        justifyContent: 'center',
        alignItems: 'start',
        paddingBottom: '128px', // Add some extra space so the FAB doesn't overlay the actions.
        gridAutoRows: 'min-content',
        gridTemplateColumns: {
          sm: '220px 220px',
          xs: '1fr',
        },
        width: {
          sm: 'auto',
        }
      }}>
        {configurations}
      </Box>
      <Fab sx={{
        position: 'absolute',
        bottom: '24px',
        right: '24px',
      }}
        color='primary'
        onClick={() => navigate(ADD_ROUTINE_URL)}>
        <AddIcon />
      </Fab>
      <Dialog
        open={deleteDialogOpen}
        onClose={() => setDeleteDialogOpen(false)}
      >
        <DialogTitle>Delete {configToDelete?.Device?.Name ?? ''}?</DialogTitle>
        <Typography
          sx={{ padding: '18px' }}
          variant="body1" >
          Deleting this routine is irreversible. Are you sure that you wish to continue?
        </Typography>
        <DialogActions>
          <Button onClick={() => onDeleteConfiguration()}>
            {!deleteProcessing ? 'DELETE' : (<CircularProgress />)}
          </Button>
          <Button onClick={() => setDeleteDialogOpen(false)}>
            CANCEL
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
