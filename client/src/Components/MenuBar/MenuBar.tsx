import {
  AppBar,
  Box,
  ClickAwayListener,
  IconButton,
  Menu,
  MenuItem,
  Toolbar,
  Typography,
} from '@mui/material';
import {
  Menu as MenuIcon,
} from '@mui/icons-material';
import {
  useNavigate,
} from 'react-router-dom';
import React, { useState } from 'react';

/**
 * The props type for the {@link MenuBar} component.
 */
export type MenuBarProps = {
  /**
   * The title to render in the navigation bar.
   */
  title: string;
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
  } = props;

  const navigate = useNavigate();
  const [menuIsOpen, setMenuIsOpen] = useState(false);
  const [anchorElement, setAnchorElement] = useState<null | HTMLElement>();

  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    setMenuIsOpen(true);
    setAnchorElement(e.currentTarget);
  };

  const handleClickAway = () => {
    setMenuIsOpen(false);
    setAnchorElement(null);
  };

  return (
    <Box sx={{ flexGrow: 1 }} >
      <AppBar position="static">
        <Toolbar>
          <ClickAwayListener onClickAway={handleClickAway}>
            <IconButton onClick={handleClick}>
              <MenuIcon />
            </IconButton>
          </ClickAwayListener>
          <Menu
            anchorEl={anchorElement}
            open={menuIsOpen}>
            <MenuItem onClick={() => navigate('/routines')}>Routines</MenuItem>
            <MenuItem onClick={() => navigate('/account')}>Account</MenuItem>
            <MenuItem onClick={() => navigate('/login')}>Login/Logout</MenuItem>
          </Menu>
          <Typography
            component="div"
            variant="h6" sx={{ flexGrow: 1 }}>
            SmartRoutine{title}
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
