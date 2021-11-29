import { instanceOfLoginDetailsBlob, LoginDetailsBlob } from './LoginState';

const ProdBaseUri = 'http://ec2-184-169-188-209.us-west-1.compute.amazonaws.com:8080';
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
 * @param endpoint The endpoint that should be queried with this request.
 * @param body The body associated with the provided endpoint.
 * @param requestType The type of request to submit. The default value is 'POST'.
 * @returns The request init value that should be associated with a fetch request.
 */
export async function FetchRequest<T extends keyof Endpoints>(
  endpoint: T,
  body: Endpoints[T],
  requestType?: 'POST' | 'GET',
) {
  const endpointURL = EndpointTargets[endpoint as EndpointTargets];
  const url = `${GetRootURL()}${endpointURL}`;

  return await fetch(url, {
    method: requestType ?? 'POST',
    body: new URLSearchParams(body),
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
  });
}

export enum EndpointTargets {
  userRead = '/user/',
  userLogin = '/user/login/',
  userCreate = '/user/create/',
  userUpdate = '/user/update/',
  userDelete = '/user/delete/',
  deviceRead = '/device/',
  deviceReadUser = '/device/user/',
  deviceReadRoutine = '/device/routine/',
  deviceCreate = '/device/create/',
  deviceUpdate = '/device/update/',
  deviceDelete = '/device/delete/',
  routineRead = '/routine/',
  routineReadUser = '/routines/user/',
  routineReadDevice = '/routines/device/',
  routineCreate = '/routine/create/',
  routineUpdate = '/routine/update/',
  routineDelete = '/routine/delete/',
  configurationCreate = '/configuration/create/',
  configurationUpdate = '/configuration/update/',
  configurationRead = '/configuration/',
  configurationReadDevice = '/configurations/device/',
  configurationReadUser = '/configurations/user/',
  configurationReadRoutine = '/configurations/routine/',
  configurationDelete = '/configuration/delete//',
}

export interface Endpoints extends Record<EndpointTargets, Record<string, string>> {
  userRead: {
    id: string
  },
  userLogin: {
    username: string,
    password: string,
  },
  userCreate: {
    username: string,
    name: string,
    password: string,
  },
  userUpdate: {
    username: string,
    name: string,
  },
  userDelete: {
    id: string,
  },
  deviceRead: {
    id: string,
  },
  deviceReadUser: {
    userid: string,
  },
  deviceReadRoutine: {
    routineid: string,
  },
  deviceCreate: {
    name: string,
    userid: string,
  },
  deviceUpdate: {
    name: string,
    deviceid: string,
  },
  deviceDelete: {
    id: string,
  },
  routineRead: {
    routineid: string,
  },
  routineReadUser: {
    userid: string,
  },
  routineReadDevice: {
    deviceid: string,
  },
  routineCreate: {
    name: string,
    userid: string,
    basealarm: string,
  },
  routineUpdate: {
    name: string,
    routineid: string,
    basealarm: string,
  },
  routineDelete: {
    id: string,
  },
  configurationCreate: {
    offset: string,
    deviceid: string,
    routineid: string,
  },
  configurationUpdate: {
    configid: string,
    offset: string,
  },
  configurationRead: {
    id: string,
  },
  configurationReadDevice: {
    deviceid: string,
  },
  configurationReadUser: {
    userid: string,
  },
  configurationReadRoutine: {
    routineid: string,
  },
  configurationDelete: {
    id: string,
  },
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
 * Parses the provided JSON input and converts that input to a {@link StoredRoutine}, if
 * possible.
 * 
 * @param jsonInput The JSON string input to evaluate.
 * @returns The parsed object, if applicable. Undefined if the conversion is not possible.
 */
 export function ParseDevice(jsonInput: string): StoredDevice | undefined {
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

  let result: StoredDevice | undefined = undefined;

  try {
    const parsedObject = JSON.parse(inputSubstr, (key, value) => {
      if (key === 'BaseAlarm') {
        return new Date(value);
      }
      return value;
    });
    if (instanceOfStoredDevice(parsedObject)) {
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
export function ParseDeviceArray(jsonInput: string): StoredDevice[] | undefined {
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

  const result: StoredDevice[] | undefined = [];

  try {
    const jsonData = JSON.parse(inputSubstr);

    if (Array.isArray(jsonData)) {
      for (let i = 0; i < jsonData.length; i++) {
        // We need to revive this differently from the base implementation.
        const reStringifiedJson = JSON.stringify(jsonData[i]);
        const device = ParseDevice(reStringifiedJson);
        if (instanceOfStoredDevice(device)) {
          result.push(device);
        }
      }
    }

  } catch (e) {
    console.error(e);
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

  result = {
    ...result,
    Configuration: result.Configuration.filter((c) => c.Id && c.Device && c.Device.Id)
  };

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
