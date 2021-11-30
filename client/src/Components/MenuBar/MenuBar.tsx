import {
  AppBar,
  Box,
  IconButton,
  styled,
  Toolbar,
  Typography,
} from '@mui/material';
import {
  DarkMode,
  LightMode,
  Menu as MenuIcon,
} from '@mui/icons-material';
import React, {
  useMemo,
} from 'react';
import {
  useColorMode,
} from '../../Utils/ColorContext';
import {
  SxProps,
  Theme,
} from '@mui/system';

/**
 * The props type for the {@link MenuBar} component.
 */
export type MenuBarProps = {
  /**
   * The title to render in the navigation bar.
   */
  title: string;
  sx?: SxProps<Theme>
  handleClickMenu: () => void;
}

/**
 * A navigation header bar that should be rendered across all views.
 * 
 * @param props See {@link MenuBarProps}.
 * @returns The component.
 */
export default function MenuBar(props: MenuBarProps) {
  const {
    title,
    sx,
    handleClickMenu,
  } = props;

  const colorState = useColorMode();
  const colorModeButton = useMemo(() => colorState?.colorMode === 'dark'
    ? (<DarkMode />)
    : (<LightMode />),
    [colorState],
  );

  const toggleColorMode = () => {
    if (colorState !== undefined) {
      colorState.colorMode === 'dark'
        ? colorState.setColorMode('light')
        : colorState.setColorMode('dark');
    }
  };

  return (
    <Box sx={{ ...sx, flexGrow: 1 }} >
      <AppBar position="static">
        <Toolbar>
          <IconButton sx={{ display: { xs: 'block', sm: 'none' } }} onClick={handleClickMenu}>
            <MenuIcon />
          </IconButton>
          <Typography
            component="div"
            variant="h6" sx={{ flexGrow: 1 }}>
            {title}
          </Typography>
          <StyledDiv>
            <IconButton onClick={toggleColorMode}>
              {colorModeButton}
            </IconButton>
          </StyledDiv>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

const StyledDiv = styled('div')``;
