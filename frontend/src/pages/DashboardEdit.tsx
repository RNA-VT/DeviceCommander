import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'

import PositionedButtonContainer from "../components/device_management/PositionedButtonContainer"

const DeviceList = () => {
  return (
    <Subscribe to={[DeviceManagement]}>
      {DeviceManagement => (
        <div>
          <PositionedButtonContainer deviceManager={DeviceManagement} />
        </div>
      )}
    </Subscribe>
  )
}

const DashboardEdit = () => (
  <Wrapper>
    <Container>
      <List>
        <DeviceList />
      </List>
    </Container>
  </Wrapper >
)

export default DashboardEdit
