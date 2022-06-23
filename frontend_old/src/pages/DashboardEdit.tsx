import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'

import PositionedButtonContainer from "../components/device_management/PositionedButtonContainer"
import ControlPanelContainer from '../containers/ControlPanelContainer'

const ControlPanel = () => {
  return (
    <Subscribe to={[DeviceManagement, ControlPanelContainer]}>
      {(DeviceManagement, controlPanelManagement) => (
        <div>
          <PositionedButtonContainer
            deviceManager={DeviceManagement}
            controlPanelManager={controlPanelManagement} />
        </div>
      )}
    </Subscribe>
  )
}

const DashboardEdit = () => (
  <Wrapper>
    <Container>
      <List>
        <ControlPanel />
      </List>
    </Container>
  </Wrapper >
)

export default DashboardEdit
