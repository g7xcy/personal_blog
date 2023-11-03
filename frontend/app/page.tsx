'use client'
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { createTheme, ThemeProvider, Grid, Box, List, createStyles, Theme, ListItem, ListItemText, Typography, Divider  } from '@mui/material'
import makeStyles from '@mui/styles/makeStyles';

const theme = createTheme();

const useStyles = makeStyles((theme: Theme) =>
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

function AppListContent(props: any) {
  const classes = useStyles(theme);
  return <List className={classes.root} {...props}> </List>;
}

function AppListItemContent(props: any) {
  const classes = useStyles(theme);
  return <ListItemText className={classes.inline} {...props}></ListItemText>
}

export default function Page() {
  const [blogs, setBlogs] = useState([]);
  const router = useRouter();

  useEffect(() => {
    // 在这里发起后端请求获取数据
    const fetchData = async (pageNumber = 1, pageSize = 5) => {
      try {
        const response = await fetch(`http://localhost/blogs/?page=${pageNumber}&size=${pageSize}`, {
          headers: {
          'Content-Type': 'application/json'
          },
      })
        const result = await response.json()
        setBlogs(result);
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    }

    fetchData()
  }, [])

  const handleItemClick = (id: number) => {
    // 使用Next.js的Router进行导航
    router.push(`http://${process.env.BASE_URL}:${process.env.BACKEND_PORT}/blog/id/${id}`);
  };

  return (
    <>
    <Box style={{ flexGrow: 1, textAlign: 'center', top: '30%', paddingTop: '2%'}}>
    <Grid container spacing={1}>
      <Grid item xs={12}>
        <h1> 欢迎来到爵爵子的Blog </h1>
      </Grid>
      <Grid item xs={12}>
      <ThemeProvider theme={theme}>
      <AppListContent style={{flexGrow: 1, paddingLeft: '20%', width: '60%'}}/>
      <ListItem alignItems="flex-start">
        {/* <ListItemAvatar>
          <Avatar alt="Remy Sharp" src="/static/images/avatar/1.jpg" />
        </ListItemAvatar> */}
          {blogs.map((blog) => (
            <AppListItemContent
            key={blog.Blog_id}
            onClick={() => handleItemClick(blog.Blog_id)}
            primary={blog.Title}
            secondary={
              <React.Fragment>
                <Typography
                  component="span"
                  variant="body2"
                  color="textPrimary"
                >
                  {blog.User}
                </Typography>
                 {" | " + blog.Content.slice(0, 20)}
              </React.Fragment>
            }
          />
          ))}
      </ListItem>
      <Divider variant="inset" component="li" />
    </ThemeProvider>
      </Grid>
    </Grid>
    </Box></>
  )
}
