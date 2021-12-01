import {
  LoadingButton,
} from '@mui/lab';
import {
  TextField,
  Typography,
} from '@mui/material';
import {
  Box,
  styled,
} from '@mui/system';
import React, {
  useEffect,
  useState,
} from 'react';
import {
  useSearchParams,
} from 'react-router-dom';
import AlertsBox from '../../Components/Containers/AlertsBox';
import {
  FetchRequest,
  ParseConfiguration,
  StoredConfiguration,
} from '../../Utils/BackendIntegration';
import {
  CONFIGURATION_ID_SEARCH_PARAM,
} from '../../Utils/CommonRouting';

/**
 * A view used to edit an existing configuration on a routine.
 * 
 * @returns The view.
 */
export default function EditConfigurationView() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [searchParams, _] = useSearchParams();
  const [configuration, setConfiguration] = useState<StoredConfiguration | undefined>(undefined);

  const [offset, setOffset] = useState(0);

  const [isSubmitting, setIsSubmitting] = useState(false);
  const [genericError, setGenericError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const configurationId = searchParams.get(CONFIGURATION_ID_SEARCH_PARAM) ?? '';

  const loadConfiguration = async () => {
    try {
      const response = await FetchRequest('configurationRead', {
        id: configurationId,
      });

      const text = await response.text();
      if (!response.ok) {
        throw text;
      }

      const config = ParseConfiguration(text);
      if (config) {
        setConfiguration(config);
        setOffset(config.Offset);
      } else {
        throw 'There was an error fetching the configuration. Please try again.';
      }
    } catch (e) {
      console.error(e);
      setGenericError(`${e}`);
    }
  };

  useEffect(() => {
    loadConfiguration();
  }, [configurationId]);

  const editConfiguration = async () => {
    try {
      const response = await FetchRequest('configurationUpdate', {
        configid: configurationId,
        offset: `${offset}`,
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

  const validateInput = (): string | undefined => {
    if (typeof offset !== 'number') {
      return 'Please enter a number for the offset value.';
    }

    return undefined;
  };

  const onSubmit: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    const validationError = validateInput();
    if (!!validationError) {
      setGenericError(validationError);
      setIsSubmitting(false);
    } else {
      try {
        await editConfiguration();
        setSuccessMessage('Successfully updated the configuration!');
      } catch (e) {
        console.error(e);
        setGenericError(`${e}`);
      } finally {
        setIsSubmitting(false);
      }
    }
  };

  return (
    <Box
      sx={{
        textAlign: 'center',
        display: 'grid',
        gridTemplateAreas: `
          "title"
          "alerts"
          "form"
        `,
      }}>
      <Typography
        sx={{
          gridArea: 'title',
        }}
        variant="h3">Edit {configuration?.Device?.Name ?? ''} Configuration</Typography>
      <AlertsBox errorMessage={genericError} successMessage={successMessage}/>
      <StyledForm
        onSubmit={onSubmit}
        sx={{
          gridArea: 'form',
          paddingTop: '18px',
          display: 'grid',
          rowGap: '12px',
          justifyContent: 'center',
          alignItems: 'center',
          gridTemplateAreas: `
            "offset"
            "submit"
          `,
        }}>
        <TextField
          value={offset}
          onChange={(e) => {
            const offsetInput = e.target.value;
            if (/^[0-9]+$/.test(offsetInput)) {
              const parsedInput = parseInt(offsetInput);
              const valueToSet = parsedInput === NaN ? 0 : parsedInput;
              setOffset(valueToSet);
            }
          }}
          label="Device Offset"
          id="device-offset"
          type="number"
          sx={{
            gridArea: 'offset'
          }} />
          <LoadingButton
            sx={{
              gridArea: 'submit',
            }}
            type="submit"
            loading={isSubmitting || !configuration}>
            <Typography variant="button">Update this configuration</Typography>
          </LoadingButton>
      </StyledForm>
    </Box>
  );
}

const StyledForm = styled('form')``;
