import {
  FormControl,
  InputLabel,
  FilledInput,
  FormHelperText,
  FormControlProps,
} from '@mui/material';
import React from 'react';

/**
 * A type mask for the react change event handler on input/text area changes.
 */
export type OnValidatedInputChange = React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;

/**
 * The props for {@link ValidatedInput}.
 */
export type ValidatedInputProps = FormControlProps & {
  /**
   * The text to render on the label.
   */
  labelText: string;

  /**
   * The id used to refer to this particular input.
   */
  labelId: string;

  /**
   * The value to render inside of the input.
   */
  value: string;

  /**
   * The error message to render when the input is invalid. If this is empty/undefined, then no
   * error will be rendered.
   */
  errorMessage: string;

  /**
   * The callback event when the input is changed.
   */
  onValueChange?: React.ChangeEventHandler<HTMLTextAreaElement | HTMLInputElement>
}

/**
 * A controlled form input that provides an easier-to-use error message input field.
 * 
 * @param props See {@link ValidatedInputProps}.
 * @returns The component.
 */
export default function ValidatedInput(props: ValidatedInputProps) {
  const {
    labelId,
    labelText,
    value,
    errorMessage,
    onValueChange,
    ...others
  } = props;

  return (
    <FormControl {...others} variant="filled" error={!!errorMessage}>
      <InputLabel htmlFor={labelId}>{labelText}</InputLabel>
      <FilledInput id={labelId} value={value} onChange={onValueChange} />
      <FormHelperText>{errorMessage}</FormHelperText>
    </FormControl>
  );
}
