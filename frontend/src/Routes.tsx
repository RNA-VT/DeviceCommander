import * as React from 'react'
import { Route, Switch } from 'react-router-dom'
import Home from './pages/Home'
import About from './pages/About'
import DeviceManagementPage from './pages/DeviceManagementPage'
import DashboardEdit from './pages/DashboardEdit'
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles'
import CssBaseline from '@material-ui/core/CssBaseline'
import { green, orange, red, grey } from '@material-ui/core/colors'

import getMuiTheme from '@material-ui/styles/'

const globalTheme = createMuiTheme({
  palette: {
    primary: {
      main: grey[900]
    },
    type: 'dark'
  },
});

const Routes = () => (
  <ThemeProvider theme={globalTheme}>
    <CssBaseline />
    <Switch>
      <Route exact path='/' component={Home} />
      <Route exact path='/about' component={About} />
      <Route exact path='/device-management' component={DeviceManagementPage} />
      <Route exact path='/dashboard/edit' component={DashboardEdit} />
    </Switch>
  </ThemeProvider>
)

export default Routes
