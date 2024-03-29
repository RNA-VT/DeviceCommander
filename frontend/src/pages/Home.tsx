import React, { useEffect } from 'react';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import { useSetRecoilState } from 'recoil';
import Dashboard from '../layouts/dashboard/Dashboard';
import { PageState } from './store';

function Home() {
  const setPageState = useSetRecoilState(PageState);

  useEffect(() => {
    setPageState({
      title: 'Home',
      index: 'home',
    });
  });

  return (
    <Dashboard>
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Grid container spacing={3}>
          {/* Chart */}
          <Grid item xs={12} md={8} lg={9}>
            <Paper
              sx={{
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                height: 240,
              }}
            >
              {/* <Chart /> */}
              <h1>Chart</h1>
            </Paper>
          </Grid>
          {/* Recent Deposits */}
          <Grid item xs={12} md={4} lg={3}>
            <Paper
              sx={{
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                height: 240,
              }}
            >
              <h1>Deposits</h1>
              {/* <Deposits /> */}
            </Paper>
          </Grid>
          {/* Recent Orders */}
          <Grid item xs={12}>
            <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
              <h1>Orders</h1>
              {/* <Orders /> */}
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </Dashboard>

  );
}

export default Home;
