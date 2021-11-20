import {
  createContext,
  useContext,
  useEffect,
  useMemo,
  useState,
} from 'react';

/**
 * The constant key used to refer to what the login state is stored as for the user.
 */
const LOGIN_STATE_KEY = 'loginstate';

/**
 * Used to determine if the given object is valid as a login details blob type.
 * 
 * @param object The object to evaluate.
 * @returns A typeguarded truthy value if the object is a valid {@link LoginDetailsBlob}.
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function instanceOfLoginDetailsBlob(object: any): object is LoginDetailsBlob {
  return (
    'Username' in object &&
    'Name' in object
  );
}

/**
 * A type used to refer to the login details for the current user.
 */
export class LoginDetailsBlob {
  /**
   * The username for the active user.
   */
  Username: string;

  /**
   * The real name for the active user.
   */
  Name: string;
}

/**
 * A class representing the authentication state for the user at any given moment.
 */
export class AuthState {
  /**
   * If true, the initial authentication pass has already happened. This is used primary for the
   * routing system so that we don't route a user before we've had the opportunity to validate their
   * stored credentials.
   */
  attemptedToGetState: boolean;

  /**
   * The login details for the current user.
   */
  loginDetails: LoginDetailsBlob | null;

  /**
   * A callback event that signs a user out of their current session.
   */
  signOut: (() => void) | null;

  /**
   * A callback event that updates the stored credentials for the user. This should actually be
   * called after the user has been logged in successfully.
   */
  signIn: ((loginDetails: LoginDetailsBlob) => void) | null;
}

/**
 * A hook used to fetch the active login state for the user. Currently, this just listens to the
 * local storage value to determine the user's login state. This could be modified to a websocket
 * subscription listener in future.
 * 
 * @returns The current login state, if any. If there is no active login undefined is returned
 * instead.
 */
function useLoginState(): AuthState | undefined {
  const [loginDetails, setLoginDetails] = useState<LoginDetailsBlob | undefined>();
  const [attemptedToGetState, setAttemptedToGetState] = useState(false);
  const storedLogin = localStorage.getItem(LOGIN_STATE_KEY);

  useEffect(() => {
    let login: LoginDetailsBlob | undefined = undefined;

    if (storedLogin !== null) {
      const parsedLogin = JSON.parse(storedLogin);
      if (instanceOfLoginDetailsBlob(parsedLogin)) {
        login = parsedLogin as LoginDetailsBlob;
      } else {
        // If our login state is invalid, purge it!
        setLoginDetails(undefined);
      }
    }

    setLoginDetails(login);

    if (!attemptedToGetState) {
      setAttemptedToGetState(true);
    }

    return () => { setLoginDetails(undefined); };
  }, [storedLogin]);

  const signOut = useMemo(() => () => {
    setLoginState(undefined);
    setLoginDetails(undefined);
  }, [setLoginDetails]);

  const signIn = useMemo(() => (loginDetails: LoginDetailsBlob) => {
    setLoginState(loginDetails);
    setLoginDetails(loginDetails);
  }, [setLoginDetails]);

  return {
    loginDetails,
    signOut,
    signIn,
    attemptedToGetState,
   };
}

/**
 * This context is used to provide the authorization state throughout the app.
 * https://usehooks.com/useAuth/
 */
const authContext = createContext<AuthState | undefined>(undefined);

/**
 * This provider is used to provide the state for the current user's authorization throughout the
 * application. This should be provided close to the root.
 * 
 * @param props Currently nothing other than children.
 * @returns The context provider.
 */
export function AuthProvider(props: React.PropsWithChildren<Record<string, unknown>>) {
  const {
    children,
  } = props;

  const auth = useLoginState();

  return (
    <authContext.Provider value={auth}>
      {children}
    </authContext.Provider>
  );
}

/**
 * This hook is used to fetch the authorization state for the active user.
 *
 * @returns The authorization state for the active user.
 */
export function useAuth(): AuthState | undefined {
  return useContext(authContext);
}

/**
 * Sets the login state for the current user.
 * 
 * @param loginState The login state for the current user. If undefined is provided, the user is
 * considered to have been 'logged out'.
 */
function setLoginState(loginState: LoginDetailsBlob | undefined) {
  if (loginState === undefined) {
    localStorage.removeItem(LOGIN_STATE_KEY);
  } else {
    const serializedState = JSON.stringify(loginState);
    localStorage.setItem(LOGIN_STATE_KEY, serializedState);
  }
}
 