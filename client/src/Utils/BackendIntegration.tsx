import { instanceOfLoginDetailsBlob, LoginDetailsBlob } from './LoginState';

const ProdBaseUri = 'http://ec2-54-153-120-74.us-west-1.compute.amazonaws.com:8080';
const LocalBaseUri = 'http://localhost:8080';

/**
 * A class representing a device as it is stored in the DB.
 */
export class StoredDevice {
  /**
   * The ID for the device.
   */
  Id: string;

  /**
   * The name of the stored device.
   */
  Name: string;

  /**
   * The ID for the user to whom the device belongs.
   */
  UserID: string;
}

/**
 * A class representing a configuration as it is stored in the DB.
 */
export class StoredConfiguration {
  /**
   * The ID for the configuration.
   */
  Id: string;

  /**
   * The ID of the associated routine.
   */
  RoutineId: string;

  /**
   * The offset value for the associated configuration.
   */
  Offset: number;

  /**
   * The device associated with the configuration.
   */
  Device: StoredDevice;
}

/**
 * A class representing a routine as it is stored in the DB.
 */
export class StoredRoutine {
  /**
   * The ID for the routine.
   */
  Id: string;

  /**
   * The name of the stored routine.
   */
  Name: string;

  /**
   * The ID for the user to whom the routine belongs.
   */
  UserId: string;

  /**
   * The time for the base alarm that is established for the routine.
   */
  BaseAlarm: Date;

  /**
   * The array of configurations associated with this routine.
   */
  Configuration: StoredConfiguration[];
}

/**
 * Gets a backend-parseable string from the given routine's basealarm field.
 * 
 * @param time The date to evaluate to get the string.
 * @returns The parsed string from the routine's basealarm field.
 */
export function GetRoutineBasealarmString(time: Date): string {
  const timeZoneOffsetHours = Math.ceil(time.getTimezoneOffset() / 60);
  const offsetString = timeZoneOffsetHours >= 10
    ? `${timeZoneOffsetHours}:00`
    : `0${timeZoneOffsetHours}:00`;

  const timeHours = time.getHours();
  const timeMinutes = time.getMinutes();

  const formattedHours = timeHours >= 10
    ? `${timeHours}`
    : `0${timeHours}`;

  const formattedMinutes = timeMinutes >= 10
    ? `${timeMinutes}`
    : `0${timeMinutes}`;

  const timeString = `${formattedHours}:${formattedMinutes}-${offsetString}`;

  return timeString;
}

/**
 * A function used to get the base URL for the backend services. Will automatically target localhost
 * or the prod env based on the current URL.
 * 
 * @returns The root URL for the backend services.
 */
export function GetRootURL(): string {
  // https://docs.amplify.aws/lib/auth/social/q/platform/js/#configure-auth-category
  const isLocalhost = Boolean(
    window.location.hostname === 'localhost' ||
    // [::1] is the IPv6 localhost address.
    window.location.hostname === '[::1]' ||
    // 127.0.0.1/8 is considered localhost for IPv4.
    window.location.hostname.match(
      /^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/
    )
  );

  return isLocalhost ? LocalBaseUri : ProdBaseUri;
}

/**
 * A generic method wrapper to send a fetch request. The BE expects form data, so this simplifies
 * some of the requests.
 * 
 * @param body The body to send. This is really just a mapped object type.
 * @param requestType The type of request to send. The default value is 'POST'.
 * @returns The request init value that should be associated with a fetch request.
 */
export function GetFetchRequest(
  body: Record<string, string>,
  requestType?: 'POST' | 'GET',
): RequestInit {
  const useRequestType = requestType ?? 'POST';
  return {
    method: useRequestType,
    body: new URLSearchParams(body),
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    }
  } as RequestInit;
}

/**
 * Gets the login URL for the current session.
 * 
 * @returns The login URL for the current session.
 */
export function GetLoginURL() {
  return `${GetRootURL()}/user/login/`;
}

/**
 * Gets the signup URL for the current session.
 * 
 * @returns The signup URL for the current session.
 */
export function GetSignupURL() {
  return `${GetRootURL()}/user/create/`;
}

/**
 * Gets the URL for creating a new routine.
 * 
 * @returns The URL for creating a new routine.
 */
export function GetCreateRoutineURL() {
  return `${GetRootURL()}/routine/create/`;
}

/**
 * Gets the URL for modifying a user's information.
 * 
 * @returns The URL for modifying a user's information.
 */
export function GetModifyUserURL() {
  return `${GetRootURL()}/user/update/`;
}

/**
 * Gets the URL for deleting a routine.
 * 
 * @returns The URL for deleting a routine.
 */
export function GetDeleteRoutineURL() {
  return `${GetRootURL()}/routine/delete/`;
}

/**
 * Gets the URL for modifying a given routine.
 * 
 * @returns The URL for modifying a routine.
 */
export function GetUpdateRoutineURL() {
  return `${GetRootURL()}/routine/update/`;
}

/**
 * Gets the URL for fetching a routine.
 * 
 * @returns The URL for fetching a routine.
 */
export function GetGetRoutineURL() {
  return `${GetRootURL()}/routine/`;
}

/**
 * Gets the routines fetch URL for the current user.
 * 
 * @returns The routines fetch URL for the current user.
 */
export function GetRoutinesFetchURL() {
  return `${GetRootURL()}/routines/user/`;
}

/**
 *
 */
export function GetRoutineDeleteURL() {
  return `${GetRootURL()}/routine/delete/`;
}

/**
 * Evaluates a login response. This resopnse may come either from the login call or from the signup
 * call.
 * 
 * @param jsonInput The response to evaluate from a fetch call either to login or to signup.
 * @returns The evaluated login response. Returns undefined if the string is an invalid response.
 */
export function ParseLoginResponse(jsonInput: string): LoginDetailsBlob | undefined {
  let openingBracket: number | undefined = undefined;
  let closingBracket: number | undefined = undefined;

  [...jsonInput].forEach((c, i) => {
    if (openingBracket === undefined) {
      openingBracket = c === '{' ? i : undefined;
    }
    // The '+1' is to adjust for the 0/1-indexing mismatch between .substring() and an item's index.
    closingBracket = c === '}' ? i + 1 : closingBracket;
  });

  const inputSubstr = jsonInput.substring(openingBracket, closingBracket);
  let result: LoginDetailsBlob | undefined = undefined;

  try {
    const parsedObject = JSON.parse(inputSubstr);
    if (instanceOfLoginDetailsBlob(parsedObject)) {
      result = parsedObject as LoginDetailsBlob;
    }
  } catch (e) {
    console.error(e);
  }
  
  return result;
}

/**
 * This type-guard function determines if the given object is convertible to a {@link StoredDevice}.
 * 
 * @param object The object to observe the typeguard.
 * @returns A typeguard declaring the object as a valid {@link StoredDevice}.
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function instanceOfStoredDevice(object: any): object is StoredDevice {
  return (
    'Id' in object &&
    'Name' in object &&
    'UserID' in object
  );
}

/**
 * This type-guard function determines if the given object is convertible to a
 * {@link StoredConfiguration}.
 * 
 * @param object The object to observe the typeguard.
 * @returns A typeguard declaring the object as a valid {@link StoredConfiguration}.
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function instanceOfStoredConfiguration(object: any): object is StoredConfiguration {
  return (
    'Id' in object &&
    'RoutineId' in object &&
    'Offset' in object &&
    'Device' in object
  );
}

/**
 * This type-guard function determines if the given object is convertible to a
 * {@link StoredRoutine}.
 * 
 * @param object The object to observe the typeguard.
 * @returns A typeguard declaring the object as a valid {@link StoredRoutine}.
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function instanceOfStoredRoutine(object: any): object is StoredRoutine {
  return (
    'Id' in object &&
    'Name' in object &&
    'UserId' in object &&
    'Configuration' in object
  );
}

/**
 * Parses the provided JSON input and converts that input to a {@link StoredDevice}, if possible.
 * 
 * @param jsonInput The JSON string input to evaluate.
 * @returns The parsed object, if applicable. Undefined if the conversion is not possible.
 */
export function ParseDevice(jsonInput: string): StoredDevice | undefined {
  let result: StoredDevice | undefined = undefined;
  const parsedObject = JSON.parse(jsonInput);

  if (instanceOfStoredDevice(parsedObject)) {
    result = parsedObject;
  }

  return result;
}

/**
 * Parses the provided JSON input and converts that input to a {@link StoredConfiguration}, if
 * possible.
 * 
 * @param jsonInput The JSON string input to evaluate.
 * @returns The parsed object, if applicable. Undefined if the conversion is not possible.
 */
export function ParseConfiguration(jsonInput: string): StoredConfiguration | undefined {
  let result: StoredConfiguration | undefined = undefined;
  const parsedObject = JSON.parse(jsonInput);

  if (instanceOfStoredConfiguration(parsedObject)) {
    result = parsedObject;
  }

  return result;
}

/**
 * Parses the provided JSON input and converts that input to a {@link StoredRoutine}, if
 * possible.
 * 
 * @param jsonInput The JSON string input to evaluate.
 * @returns The parsed object, if applicable. Undefined if the conversion is not possible.
 */
export function ParseRoutine(jsonInput: string): StoredRoutine | undefined {
  let openingBracket: number | undefined = undefined;
  let closingBracket: number | undefined = undefined;

  [...jsonInput].forEach((c, i) => {
    if (openingBracket === undefined) {
      openingBracket = c === '{' ? i : undefined;
    }
    // The '+1' is to adjust for the 0/1-indexing mismatch between .substring() and an item's index.
    closingBracket = c === '}' ? i + 1 : closingBracket;
  });

  const inputSubstr = jsonInput.substring(openingBracket, closingBracket);

  let result: StoredRoutine | undefined = undefined;

  try {
    const parsedObject = JSON.parse(inputSubstr, (key, value) => {
      if (key === 'BaseAlarm') {
        return new Date(value);
      }
      return value;
    });
    if (instanceOfStoredRoutine(parsedObject)) {
      result = parsedObject;
    }
  } catch (e) {
    console.error(e);
  }

  return result;
}

/**
 * Parses the provided JSON input and converts that input to a {@link StoredRoutine}[], if
 * possible.
 * 
 * @param jsonInput The JSON string input to evaluate.
 * @returns The parsed object, if applicable. Undefined if the conversion is not possible.
 */
export function ParseRoutineArray(jsonInput: string): StoredRoutine[] | undefined {
  let openingBracket: number | undefined = undefined;
  let closingBracket: number | undefined = undefined;

  [...jsonInput].forEach((c, i) => {
    if (openingBracket === undefined) {
      openingBracket = c === '[' ? i : undefined;
    }
    // The '+1' is to adjust for the 0/1-indexing mismatch between .substring() and an item's index.
    closingBracket = c === ']' ? i + 1 : closingBracket;
  });
  
  const inputSubstr = jsonInput.substring(openingBracket, closingBracket);

  const result: StoredRoutine[] | undefined = [];

  try {
    const jsonData = JSON.parse(inputSubstr);

    if (Array.isArray(jsonData)) {
      for (let i = 0; i < jsonData.length; i++) {
        // We need to revive this differently from the base implementation.
        const reStringifiedJson = JSON.stringify(jsonData[i]);
        const routine = ParseRoutine(reStringifiedJson);
        if (instanceOfStoredRoutine(routine)) {
          result.push(routine);
        }
      }
    }

  } catch (e) {
    console.error(e);
  }

  return result;
}
