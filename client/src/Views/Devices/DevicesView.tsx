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
} from 'react-router';
import {
  FetchRequest,
  ParseDeviceArray,
  StoredDevice,
} from '../../Utils/BackendIntegration';
import {
  useAuth,
} from '../../Utils/LoginState';
import AddIcon from '@mui/icons-material/Add';
import {
  ADD_DEVICE_URL,
  DEVICE_ID_SEARCH_PARAM,
  EDIT_DEVICE_URL,
} from '../../Utils/CommonRouting';
import DeviceCard from '../../Components/Devices/DeviceCard';
import {
  LoadingButton,
} from '@mui/lab';
import {
  LoadingCardBox,
} from '../../Components/Containers/CardBox';
import AlertsBox from '../../Components/Containers/AlertsBox';

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
  const [isFetchingDevices, setIsFetchingDevices] = useState(true);

  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [deviceToDelete, setDeviceToDelete] = useState<StoredDevice | undefined>(undefined);
  const [deleteProcessing, setDeleteProcessing] = useState(false);

  const [genericError, setGenericError] = useState('');

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
      setGenericError(e);
    }
  };

  const deleteDevice = async () => {
    try {
      const response = await FetchRequest('deviceDelete', {
        id: deviceToDelete.Id,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }
    } catch (e) {
      console.error(e);
      setGenericError(e);
    }
  };

  const handleEditDevice = (device: StoredDevice): void => {
    navigate(`${EDIT_DEVICE_URL}?${DEVICE_ID_SEARCH_PARAM}=${device?.Id ?? ''}`);
  };

  const handleDeleteDevice = (device: StoredDevice): void => {
    setDeleteDialogOpen(true);
    setDeviceToDelete(device);
  };

  const onDeleteDevice = async () => {
    setDeleteProcessing(true);
    await deleteDevice();
    setDeleteProcessing(false);

    setDeleteDialogOpen(false);
    await fetchDevices();
  };

  useEffect(() => {
    fetchDevices().then(() => {
      setIsFetchingDevices(false);
    });
  }, []);

  const deviceCards = useMemo(() => devices.map((d) => (
    <DeviceCard
      device={d}
      onDeleteDevice={handleDeleteDevice}
      onEditDevice={handleEditDevice}
    />
  )), [devices]);

  return (
    <Box
      sx={{
        display: 'grid',
        gridTemplateAreas: `
          "title"
          "alerts-box"
          "devices"
        `,
        textAlign: 'center',
      }}>
      <Typography sx={{ gridArea: 'title' }} variant="h3" >
        Devices
      </Typography>
      <AlertsBox errorMessage={genericError} />
      <LoadingCardBox isLoading={isFetchingDevices} sx={{ gridArea: 'devices' }}>
        {deviceCards}
      </LoadingCardBox>
      <Fab onClick={() => navigate(ADD_DEVICE_URL)}>
        <AddIcon />
      </Fab>
      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Delete?</DialogTitle>
        <Typography>
          Deleting this device is irreversible. Are you sure that you wish to continue?
        </Typography>
        <DialogActions>
          <LoadingButton onClick={() => onDeleteDevice()} loading={deleteProcessing}>
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
