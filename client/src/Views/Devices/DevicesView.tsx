import { Alert, Box, Button, Dialog, DialogActions, DialogTitle, Fab, Typography } from '@mui/material';
import React, { useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router';
import { FetchRequest, ParseDeviceArray, StoredDevice } from '../../Utils/BackendIntegration';
import { useAuth } from '../../Utils/LoginState';
import AddIcon from '@mui/icons-material/Add';
import { ADD_DEVICE_URL, DEVICE_ID_SEARCH_PARAM, EDIT_DEVICE_URL } from '../../Utils/CommonRouting';
import DeviceCard from '../../Components/Devices/DeviceCard';

/**
 * The props for the {@link DevicesView}.
 */
export type DevicesViewProps = {
}

/**
 * A view used to render information about a single device.
 * 
 * @returns The view.
 */
export default function DevicesView() {
  const [devices, setDevices] = useState<StoredDevice[]>([]);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deviceToDelete, setDeviceToDelete] = useState<StoredDevice | undefined>(undefined);
  const [errorMessage, setErrorMessage] = useState('');

  const navigate = useNavigate();

  const authState = useAuth();
  const loginDetails = authState?.loginDetails;

  const fetchDevices = async () => {
    try {
      const response = await FetchRequest('deviceReadUser', {
        userid: loginDetails.Username ?? '',
      });

      const text = await response.text();

      if (response.ok) {
        const devicesData = await ParseDeviceArray(text);
        if (devicesData !== undefined) {
          setDevices(devicesData);
        }
      } else {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setErrorMessage(e);
    }
  };

  const handleEditDevice = (device: StoredDevice): void => {
    navigate(`${EDIT_DEVICE_URL}?${DEVICE_ID_SEARCH_PARAM}=${device?.Id ?? ''}`);
  };

  const handleDeleteDevice = (device: StoredDevice): void => {
    setDeleteDialogOpen(true);
    setDeviceToDelete(device);
  };

  useEffect(() => {
    fetchDevices();
  }, []);

  const deviceCards = useMemo(() => devices.map((d) => (
    <DeviceCard
      device={d}
      onDeleteDevice={handleDeleteDevice}
      onEditDevice={handleEditDevice}
    />
  )), [devices]);

  return (
    <Box sx={{
      display: 'grid',
      gridTemplateAreas: `
        "title"
        "devices"
      `,
      rowGap: '48px',
      textAlign: 'center',
    }}>
      <Typography sx={{ gridArea: 'title' }} variant="h3">
        Devices
      </Typography>
      <Alert sx={{ visibility: !!errorMessage ? 'visible' : 'hidden' }} severity="error" />
      <Box sx={{
        gridArea: 'devices',
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
        {deviceCards}
      </Box>
      <Fab color="primary"
        onClick={() => navigate(ADD_DEVICE_URL)}
        sx={{
          position: 'absolute',
          bottom: '24px',
          right: '24px',
        }}>
        <AddIcon />
      </Fab>
      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Delete?</DialogTitle>
        <Typography>
          Deleting this device is irreversible. Are you sure that you wish to continue?
        </Typography>
        <DialogActions>
          <Button>
            DELETE
          </Button>
          <Button onClick={() => setDeleteDialogOpen(false)}>
            CANCEL
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}
