import { PaletteMode } from '@mui/material';
import React, {
  createContext,
  useContext,
} from 'react';

/**
 * The constant key used to refer to what the color state (theme) is stored as for the user.
 */
const COLOR_STATE_KEY = 'colorstate';

/**
 * A class representing the stored color state for the user at any given moment.
 */
export class ColorState {
  /**
   * The color mode that is currently applied for the user.
   */
  colorMode: PaletteMode;

  /**
   * A callback event that changes the color theme for the user's browser.
   */
  setColorMode: (paletteMode: PaletteMode) => void;
}

/**
 * A context used to provide theming/palette info throughout the application.
 */
const colorContext = createContext<ColorState>(undefined);

/**
 * Props for the {@link ColorProvider} component.
 */
type ColorProviderProps = {
  /**
   * The current state of the theming for the system.
   */
  colorState: ColorState;
}

/**
 * This provider is used to pass theming information throughout the system. This should be provided
 * close to the root, ideally where the theme provider lives.
 * 
 * @param props See {@link ColorProviderProps}.
 * @returns The provider.
 */
export function ColorProvider(props: React.PropsWithChildren<ColorProviderProps>) {
  const {
    children,
    colorState,
  } = props;

  return(
    <colorContext.Provider value={colorState}>
      {children}
    </colorContext.Provider>
  );
}

/**
 * A hook used to fetch the current color theming for the active user.
 * 
 * @returns The current color theming for the active user.
 */
export function useColorMode(): ColorState {
  return useContext(colorContext);
}

/**
 * A function used to store the color preference for the user. This should only be used by the same
 * component that is providing the color theming via its provider.
 * 
 * @param colorMode The color mode to store.
 */
export function storeColorPreference(colorMode: PaletteMode): void {
  localStorage.setItem(COLOR_STATE_KEY, colorMode);
}

/**
 * A funciton used to fetch the color preference for the user. This should only be used by the same
 * component that is providing the color theming via its provider.
 * 
 * @returns The stored color preference.
 */
export function fetchColorPreference(): PaletteMode {
  let preference = localStorage.getItem(COLOR_STATE_KEY) as PaletteMode;

  if (preference === null) {
    const prefersDarkTheme = matchMedia('(prefers-color-scheme: dark)');
    preference = prefersDarkTheme
      ? 'dark'
      : 'light';
  }

  return preference;
}
