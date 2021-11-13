import { useEffect, useState } from "react";

const LOGIN_STATE_KEY = 'loginstate';

function instanceOfLoginDetailsBlob(object: any): object is LoginDetailsBlob {
  return (
    'username' in object &&
    'name' in object &&
    'userid' in object
  );
}

export class LoginDetailsBlob {
  username: string;
  name: string;
  userid: string;
}

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
  }, [storedLogin])

  return loginDetails;
}

export function setLoginState(loginState: LoginDetailsBlob | undefined) {
  if (loginState === undefined) {
    localStorage.removeItem(LOGIN_STATE_KEY);
  } else {
    const serializedState = JSON.stringify(loginState);
    localStorage.setItem(LOGIN_STATE_KEY, serializedState);
  }
}
 