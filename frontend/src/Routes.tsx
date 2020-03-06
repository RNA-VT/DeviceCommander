import * as React from 'react'
import { Route, Switch } from 'react-router-dom'
import Home from './pages/Home'
import About from './pages/About'
import DeviceManagementPage from './pages/DeviceManagementPage'
import DashboardEdit from './pages/DashboardEdit'

const Routes = () => (
  <div>
    <Switch>
      <Route exact path='/' component={Home} />
      <Route exact path='/about' component={About} />
      <Route exact path='/device-management' component={DeviceManagementPage} />
      <Route exact path='/dashboard/edit' component={DashboardEdit} />
    </Switch>
  </div>
)

export default Routes
