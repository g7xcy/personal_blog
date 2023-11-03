'use client'
import * as React from 'react';
import { styled, Box, TextField, InputAdornment, Grid, Button, IconButton, Snackbar, Alert, SnackbarCloseReason, Theme, createStyles } from '@mui/material';
import { AccountCircle, Password, Visibility, VisibilityOff } from '@mui//icons-material';

const useStyles = styled((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
      maxWidth: '36ch',
      backgroundColor: theme.palette.background.paper,
    },
    inline: {
      display: 'inline',
    },
  }),
);

export default function BasicGrid() {
    const classes = useStyles();
    const [values, setValues] = React.useState({email: '', password: ''})
    const [showPassword, setShowPassword] = React.useState(false)
    const [open, setOpen] = React.useState(false);
    const [isError, setIsError] = React.useState(false)
    const [alertMsg, setAlertMsg] = React.useState('');
    const handleClickShowPassword = () => {
        setShowPassword(!showPassword)
    }
    const handleMouseDownPassword = (event: React.MouseEvent) => {
        event.preventDefault();
      }
    const handleEmailChange = (email: string) => {
        const nextValues = JSON.parse(JSON.stringify(values))
        nextValues.email = email
        setValues(nextValues)
    }
    const handlePasswordChange = (password: string) => {
        const nextValues = JSON.parse(JSON.stringify(values))
        nextValues.password = password
        setValues(nextValues)
    }
    const handleClose = (event?: Event | React.SyntheticEvent<any, Event>, reason?: SnackbarCloseReason) => {
        if (reason === 'clickaway') {
          return;
        }
        setOpen(false);
    }
    async function handleOnClick(values: {email: string, password: string}) {
        try {
            const response = await fetch(`http://${process.env.BASE_URL}:${process.env.BACKEND_PORT}/login`, {
                method: "POST",
                headers: {
                'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Email: values.email,
                    Password: values.password
                })
            })
            const data = await response.json()
            if (data.status !== 200) {
                setAlertMsg("登陆失败: "+data)
                setIsError(true)
                setOpen(true)
                return
            }
            setAlertMsg("登陆成功")
            setIsError(false)
            setOpen(true)
        } catch(error) {
            setAlertMsg("登陆失败")
            setIsError(true)
            setOpen(true)
            console.log(error)
        }
    }
  return (
    <>
    <Box className={classes.root} style={{ flexGrow: 1, textAlign: 'center', top: '30%', paddingTop: '10%'}}>
      <Grid container spacing={1}>
        <Grid item xs={12}>
            <TextField 
                style={{width: '20%',}}
                color="secondary"
                id="standard-required" 
                label="email" 
                variant="standard" 
                onChange={ (v) => handleEmailChange(v.target.value) }
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                        <AccountCircle />
                        </InputAdornment>
                    ),
                }}/>
        </Grid>
        <Grid item xs={12}>
            <TextField 
                style={{width: '20%'}}
                color="secondary"
                id="standard-required" 
                label="password" 
                type={showPassword ? 'text' : 'password'} 
                variant="standard" 
                onChange={ (v) => handlePasswordChange(v.target.value) }
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                        <Password />
                        </InputAdornment>
                    ),
                    endAdornment: (
                        <InputAdornment position="end">
                          <IconButton
                            aria-label="toggle password visibility"
                            onClick={handleClickShowPassword}
                            onMouseDown={handleMouseDownPassword}
                          >
                            {showPassword ? <Visibility /> : <VisibilityOff />}
                          </IconButton>
                        </InputAdornment>
                    )
                }}
                />
        </Grid>
        <Grid item xs={12}>
            <Button 
                variant="contained" 
                onClick={ () => handleOnClick(values) }
                style={{top: '30%'}}
            >
                登录
            </Button>
        </Grid>
      </Grid>
    </Box>
    <Snackbar
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'center',
        }}
        open={open}
        autoHideDuration={3000}
        onClose={handleClose}
      >
      <Alert onClose={handleClose} severity={isError ? 'error' : 'success'}> {alertMsg} </Alert>
    </Snackbar>
    </>
  );
}