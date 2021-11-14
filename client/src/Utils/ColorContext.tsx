import { PaletteMode } from '@mui/material';
import React, { createContext, useContext, useState } from 'react';

const COLOR_STATE_KEY = 'colorstate';

export class ColorState {
  colorMode: PaletteMode;
  setColorMode: (paletteMode: PaletteMode) => void;
}

const colorContext = createContext<ColorState>(undefined);

type ColorProviderProps = {
  colorState: ColorState;
}

/**
 * @param props
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
 *
 */
export function useColorMode(): ColorState {
  return useContext(colorContext);
}

/**
 * @param colorMode
 */
export function storeColorPreference(colorMode: PaletteMode): void {
  localStorage.setItem(COLOR_STATE_KEY, colorMode);
}

/**
 *
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
