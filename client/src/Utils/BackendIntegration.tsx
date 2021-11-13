import { LoginDetailsBlob } from './LoginState';

const ProdBaseUri = 'http://localhost:8080';
const LocalBaseUri = 'http://localhost:8080';

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
 * Gets the login URL for the current session.
 * 
 * @returns The login URL for the current session.
 */
export function GetLoginURL() {
  return `${GetRootURL()}/user/`;
}

/**
 * Gets the signup URL for the current session.
 * 
 * @returns The signup URL for the current session.
 */
export function GetSignupURL() {
  return `${GetRootURL()}/create/user`;
}

/**
 * Evaluates a login response. This resopnse may come either from the login call or from the signup
 * call.
 * 
 * @param response The response to evaluate from a fetch call either to login or to signup.
 * @returns The evaluated login response. Returns undefined if the string is an invalid response.
 */
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
