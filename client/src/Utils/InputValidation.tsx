const ValidUserNameChars = /^[0-9a-zA-Z]+$/;
const ValidRoutineNameChars = /^[0-9a-zA-Z ]+$/;
const ValidDeviceNameChars = /^[0-9a-zA-Z ]+$/;

/**
 * A function used to provide common username input validation.
 * 
 * @param username The username to validate.
 * @returns An error message, if any.
 */
export function ValidateUsername(username: string): string | undefined {
  if (username.length <= 1) {
    return 'Please enter a username at least 1 character long.';
  }

  if (!username.match(ValidUserNameChars)) {
    return 'Please only use letters and numbers in your username.';
  }

  return undefined;
}

/**
 * A function used to provoide common name input validation.
 * 
 * @param name The name to validate.
 * @returns An error message, if any.
 */
export function ValidateName(name: string): string | undefined {
  if (name.length <= 1) {
    return 'Please enter a name at least 1 character long.';
  }

  return undefined;
}

/**
 * A function used to provoide common name password validation.
 * 
 * @param password The password to validate.
 * @returns An error message, if any.
 */
export function ValidatePassword(password: string): string | undefined {
  if (password.length <= 1) {
    return 'Please enter a password at least 1 character long.';
  }

  return undefined;
}

/**
 * A function used to provoide common routine name input validation.
 * 
 * @param routineName The routine name to validate.
 * @returns An error message, if any.
 */
export function ValidateRoutineName(routineName: string): string | undefined {
  if (routineName.length <= 1) {
    return 'Please enter a name at least 1 character long.';
  }

  if (!routineName.match(ValidRoutineNameChars)) {
    return 'Please only use alphanumeric characters and spaces in your routine name.';
  }

  return undefined;
}

/**
 * A function used to provoide common name device name validation.
 * 
 * @param deviceName The device name to validate.
 * @returns An error message, if any.
 */
export function ValidateDeviceName(deviceName: string): string | undefined {
  if (deviceName.length <= 1) {
    return 'Please enter a name at least 1 character long.';
  }

  if (!deviceName.match(ValidDeviceNameChars)) {
    return 'Please only use alphanumeric characters and spaces in your device name.';
  }

  return undefined;
}
