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

export type MenuBarProps = {
  title: string;
}

export default function MenuBar(props: MenuBarProps) {
  const {
    title,
  } = props;

  const navigate = useNavigate();
  const [menuIsOpen, setMenuIsOpen] = useState(false);

  return (
    <Box sx={{ flexGrow: 1 }} >
      <AppBar position="static">
        <Toolbar>
          <ClickAwayListener onClickAway={() => { setMenuIsOpen(false) }}>
            <IconButton onClick={() => setMenuIsOpen(true)}>
              <MenuIcon />
            </IconButton>
          </ClickAwayListener>
          <Menu open={menuIsOpen}>
            <MenuItem onClick={() => navigate("/routines")}>Routines</MenuItem>
            <MenuItem onClick={() => navigate("/account")}>Account</MenuItem>
          </Menu>
          <Typography
            component="div"
            variant="h6" sx={{ flexGrow: 1 }}>
            SmartRoutine{title}
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
  )
}
