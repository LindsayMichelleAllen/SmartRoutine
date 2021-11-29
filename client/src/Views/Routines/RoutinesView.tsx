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
  FetchRequest,
  ParseRoutineArray,
  StoredRoutine,
} from '../../Utils/BackendIntegration';
import AddIcon from '@mui/icons-material/Add';
import {
  useAuth,
} from '../../Utils/LoginState';
import { useNavigate } from 'react-router';
import {
  ADD_ROUTINE_URL,
  EDIT_ROUTINE_URL,
  ROUTINE_ID_SEARCH_PARAM,
  VIEW_ROUTINE_URL,
} from '../../Utils/CommonRouting';
import RoutineCard from '../../Components/Routines/RoutineCard';
import {
  LoadingButton,
} from '@mui/lab';
import { LoadingCardBox } from '../../Components/Containers/CardBox';

/**
 * The view used to describe the available routines for the user.
 * 
 * @returns The view.
 */
export default function RoutinesView() {
  const [routines, setRoutines] = useState<StoredRoutine[]>([]);
  const [isFetchingRoutines, setIsFetchingRoutines] = useState(true);
  const [routineToDelete, setRoutineToDelete] = useState<StoredRoutine | undefined>(undefined);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deleteProcessing, setDeleteProcessing] = useState(false);

  const navigate = useNavigate();
  const authState = useAuth();

  const loginDetails = authState?.loginDetails;

  const fetchRoutines = async () => {
    try {
      const response = await FetchRequest('routineReadUser', {
        userid: loginDetails.Username ?? '',
      });

      const text = await response.text();

      if (response.ok) {
        const routinesData = await ParseRoutineArray(text);
        if (routinesData !== undefined) {
          setRoutines(routinesData);
        }
      } else {
        throw text;
      }
    } catch (e) {
      console.error(e);
    }
  };

  const deleteRoutine = async () => {
    try {
      const response = await FetchRequest('routineDelete', {
        id: routineToDelete.Id,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }
    } catch (e) {
      console.error(e);
    }
  };

  const handleEditRoutine = (routine: StoredRoutine): void => {
    navigate(`${EDIT_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routine.Id}`);
  };

  const handleDeleteRoutine = (routine: StoredRoutine): void => {
    setDeleteDialogOpen(true);
    setRoutineToDelete(routine);
  };

  const handleViewRoutine = (routine: StoredRoutine): void => {
    navigate(`${VIEW_ROUTINE_URL}?${ROUTINE_ID_SEARCH_PARAM}=${routine.Id}`);
  };

  const onDeleteRoutine = async () => {
    setDeleteProcessing(true);
    await deleteRoutine();
    setDeleteProcessing(false);

    setDeleteDialogOpen(false);
    await fetchRoutines();
  };

  useEffect(() => {
    if (loginDetails !== undefined) {
      fetchRoutines().then(() => {
        setIsFetchingRoutines(false);
      });
    }
  }, [loginDetails]);

  console.log(routines);

  const routineCards = useMemo(() => routines.map((r) => (
    <RoutineCard
      key={r.Id}
      routine={r}
      onEditRoutine={handleEditRoutine}
      onDeleteRoutine={handleDeleteRoutine}
      onViewRoutine={handleViewRoutine}
    />
  )), [routines]);

  return (
    <Box sx={{
      display: 'grid',
      gridTemplateAreas: `
        "title"
        "routines"
      `,
      gridTemplateRows: 'min-content 1fr',
      rowGap: '48px',
      textAlign: 'center',
    }}>
      <Typography
        sx={{ gridArea: 'title' }}
        variant="h3">
        Routines
      </Typography>
      <LoadingCardBox isLoading={isFetchingRoutines} sx={{ gridArea: 'routines' }}>
        {routineCards}
      </LoadingCardBox>
      <Fab onClick={() => navigate(ADD_ROUTINE_URL)}>
        <AddIcon />
      </Fab>
      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Delete {routineToDelete?.Name ?? ''}?</DialogTitle>
        <Typography
          sx={{ padding: '18px' }}
          variant="body1" >
          Deleting this routine is irreversible. Are you sure that you wish to continue?
        </Typography>
        <DialogActions>
          <LoadingButton onClick={() => onDeleteRoutine()} loading={deleteProcessing}>
            <Typography variant="button">Delete</Typography>
          </LoadingButton>
          <Button disabled={deleteProcessing} onClick={() => setDeleteDialogOpen(false)}>
            CANCEL
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
