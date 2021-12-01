import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogTitle,
  Fab,
  Typography,
} from '@mui/material';
import React, {
  useEffect,
  useMemo,
  useState,
} from 'react';
import {
  useNavigate,
  useSearchParams,
} from 'react-router-dom';
import {
  FetchRequest,
  ParseRoutine,
  StoredConfiguration,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import {
  ADD_DEVICE_TO_ROUTINE_URL,
  CONFIGURATION_ID_SEARCH_PARAM,
  EDIT_CONFIGURATION_URL,
  EDIT_ROUTINE_URL,
  ROUTINE_ID_SEARCH_PARAM,
} from '../../Utils/CommonRouting';
import ConfigurationCard from '../../Components/Configurations/ConfigurationCard';
import AddIcon from '@mui/icons-material/Add';
import RoutineCard from '../../Components/Routines/RoutineCard';
import { LoadingButton } from '@mui/lab';
import { LoadingCardBox } from '../../Components/Containers/CardBox';
import CardSkeleton from '../../Components/Skeletons/CardSkeleton';
import AlertsBox from '../../Components/Containers/AlertsBox';

/**
 * A view used to render details and controls for a single routine.
 * 
 * @returns The view.
 */
export default function SingleRoutineView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [isFetchingRoutine, setIsFetchingRoutine] = useState(true);
  const [routine, setRoutine] = useState<StoredRoutine | undefined>(undefined);
  const [genericError, setGenericError] = useState('');
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
      }
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
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
      setGenericError(`${e}`);
    }
  };

  useEffect(() => {
    loadRoutine().then(() => {
      setIsFetchingRoutine(false);
    });
  }, [routineId]);

  const handleDeleteConfiguration = (configuration: StoredConfiguration) => {
    setDeleteDialogOpen(true);
    setConfigToDelete(configuration);
  };

  const handleEditConfiguration = (configuration: StoredConfiguration) => {
    navigate(`${EDIT_CONFIGURATION_URL}?${CONFIGURATION_ID_SEARCH_PARAM}=${configuration.Id}`);
  };

  const onDeleteConfiguration = async () => {
    setDeleteProcessing(true);
    await deleteConfig();
    setDeleteProcessing(false);

    setDeleteDialogOpen(false);
    await loadRoutine();
  };

  const handleEditRoutine = (routine: StoredRoutine): void => {
    navigate(`${EDIT_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routine.Id}`);
  };

  const configurations = useMemo(() => routine?.Configuration?.map((c) => (
    <ConfigurationCard
      key={c.Id}
      configuration={c}
      onDeleteConfiguration={handleDeleteConfiguration}
      onEditConfiguration={handleEditConfiguration}
    />
  )), [routine?.Configuration]);

  return (
    <Box sx={{ textAlign: 'center' }}>
      <Box sx={{
        display: 'grid',
        rowGap: '18px',
        gridTemplateAreas: `
          "routineDetails"
          "configurationsTitle"
          "alerts"
          "configurationsList"
        `,
      }}>
        {
          isFetchingRoutine
            ? (<CardSkeleton />)
            : (
              <RoutineCard
                routine={routine}
                onEditRoutine={handleEditRoutine}
              />
            )
        }
        <Typography
          sx={{
            gridArea: 'configurationsTitle',
          }}
          variant="h4">
          Devices
        </Typography>
        <AlertsBox errorMessage={genericError}/>
        <LoadingCardBox
          sx={{
            gridArea: 'configurationsList',
          }}
          isLoading={false}>
          {configurations}
        </LoadingCardBox>
      </Box>
      <Fab onClick={() => navigate(`${ADD_DEVICE_TO_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routineId}`)}>
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
          Deleting this device is irreversible. Are you sure that you wish to continue?
        </Typography>
        <DialogActions>
          <LoadingButton onClick={() => onDeleteConfiguration()} loading={deleteProcessing}>
            <Typography variant="button">Delete</Typography>
          </LoadingButton>
          <Button onClick={() => setDeleteDialogOpen(false)} disabled={deleteProcessing}>
            CANCEL
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
