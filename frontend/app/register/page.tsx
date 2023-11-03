'use client'
import * as React from 'react';
import { Box, TextField, InputAdornment, Grid, Button, IconButton, Snackbar , Alert, SnackbarCloseReason } from '@mui/material';
import { AccountCircle, Password, Email, Visibility, VisibilityOff } from '@mui//icons-material';

export default function BasicGrid() {
    const [values, setValues] = React.useState({username: '', email: '', password: ''})
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
    const handleUsernameChange = (username: string) => {
        const nextValues = JSON.parse(JSON.stringify(values))
        nextValues.username = username
        setValues(nextValues)
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
    async function handleOnClick(values: {username: string, email: string, password: string}) {
        try {
            const response = await fetch(`http://${process.env.BASE_URL}:${process.env.BACKEND_PORT}/register`, {
                method: "POST",
                headers: {
                'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    Name: values.username,
                    Email: values.email,
                    Password: values.password
                })
            })
            const data = await response.json()
            if (data.status !== 200) {
                setAlertMsg("注册失败: "+data)
                setIsError(true)
                setOpen(true)
                return
            }
            setAlertMsg("注册成功")
            setIsError(false)
            setOpen(true)
        } catch(error) {
            setAlertMsg("注册失败")
            setIsError(true)
            setOpen(true)
            console.log(error)
        }
    }
  return (
    <>
    <Box style={{ flexGrow: 1, textAlign: 'center', top: '30%',}}>
      <Grid container spacing={1}>
      <Grid item xs={12}>
            <TextField 
                style={{width: '20%', paddingTop: '10%'}}
                color="secondary"
                id="standard-required" 
                label="username" 
                variant="standard" 
                onChange={ (v) => handleUsernameChange(v.target.value) }
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
                label="email" 
                variant="standard" 
                onChange={ (v) => handleEmailChange(v.target.value) }
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                        <Email />
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
                注册
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