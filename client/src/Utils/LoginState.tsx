import { useEffect, useState } from 'react';

const LOGIN_STATE_KEY = 'loginstate';

/**
 * Used to determine if the given object is valid as a login details blob type.
 * 
 * @param object The object to evaluate.
 * @returns A typeguarded truthy value if the object is a valid {@link LoginDetailsBlob}.
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function instanceOfLoginDetailsBlob(object: any): object is LoginDetailsBlob {
  return (
    'username' in object &&
    'name' in object &&
    'userid' in object
  );
}

/**
 * A type used to refer to the login details for the current user.
 */
export class LoginDetailsBlob {
  /**
   * The username for the active user.
   */
  username: string;

  /**
   * The real name for the active user.
   */
  name: string;

  /**
   * The user ID for the active user.
   */
  userid: string;
}

/**
 * A hook used to fetch the active login state for the user. Currently, this just listens to the
 * local storage value to determine the user's login state. This could be modified to a websocket
 * subscription listener in future.
 * 
 * @returns The current login state, if any. If there is no active login undefined is returned
 * instead.
 */
export function useLoginState(): LoginDetailsBlob | undefined {
  const [loginDetails, setLoginDetails] = useState<LoginDetailsBlob | undefined>(undefined);
  const storedLogin = localStorage.getItem(LOGIN_STATE_KEY);

  useEffect(() => {
    let login: LoginDetailsBlob | undefined = undefined;

    if (storedLogin !== null) {
      const parsedLogin = JSON.parse(storedLogin);
      if (instanceOfLoginDetailsBlob(parsedLogin)) {
        login = parsedLogin as LoginDetailsBlob;
      }
    }

    setLoginDetails(login);
  }, [storedLogin]);

  return loginDetails;
}

/**
 * Sets the login state for the current user.
 * 
 * @param loginState The login state for the current user. If undefined is provided, the user is
 * considered to have been 'logged out'.
 */
export function setLoginState(loginState: LoginDetailsBlob | undefined) {
  if (loginState === undefined) {
    localStorage.removeItem(LOGIN_STATE_KEY);
  } else {
    const serializedState = JSON.stringify(loginState);
    localStorage.setItem(LOGIN_STATE_KEY, serializedState);
  }
}
 