import { LoginDetailsBlob } from "./LoginState";

export const ProdBaseUri = 'http://localhost:8080';
export const LocalBaseUri = 'http://localhost:8080';

// https://docs.amplify.aws/lib/auth/social/q/platform/js/#configure-auth-category
export function GetRootURL(): string {
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

export function GetLoginURL() {
  return `${GetRootURL()}/user/`
}

export function ParseLoginResponse(response: string): LoginDetailsBlob | undefined {
  const values = response.split(', ');
  if (values.length >= 3) {

    return {
      username: values[0],
      name: values[1],
      // the UUID is returning with some extra text. Truncate it here.
      userid: values[2].split('%')[0],
    };
  }

  return undefined;
}
