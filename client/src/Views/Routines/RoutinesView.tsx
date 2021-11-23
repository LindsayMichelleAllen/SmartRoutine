import {
  Box,
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogTitle,
  Fab,
  Typography,
} from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
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
import { ADD_ROUTINE_URL, EDIT_ROUTINE_URL, ROUTINE_ID_SEARCH_PARAM, VIEW_ROUTINE_URL } from '../../Utils/CommonRouting';
import RoutineCard from '../../Components/Routines/RoutineCard';

/**
 * The view used to describe the available routines for the user.
 * 
 * @returns The view.
 */
export default function RoutinesView() {
  const [routines, setRoutines] = useState<StoredRoutine[]>([]);
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
      fetchRoutines();
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
      height: '100%',
      width: '100%',
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
      <Box sx={{
        gridArea: 'routines',
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
        {routineCards}
      </Box>
      <Fab sx={{ position: 'absolute', bottom: '24px', right: '24px' }}
        color='primary'
        onClick={() => navigate(ADD_ROUTINE_URL)}>
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
          <Button onClick={() => onDeleteRoutine()}>
            {
              !deleteProcessing ? 'DELETE' : (<CircularProgress />)
            }
          </Button>
          <Button onClick={() => setDeleteDialogOpen(false)}>
            CANCEL
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
