import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'

import PositionedButtonContainer from "../components/device_management/PositionedButtonContainer"

const DashboardEdit = () => (
  <Wrapper>
    <PageHeader />
    <Container maxWidth="md">
      <List>
        <Subscribe to={[DeviceManagement]}>
          {deviceManager => {
            return (
              <div>
                <PositionedButtonContainer
                  deviceManager={deviceManager} />
              </div>
            )
          }}
        </Subscribe>
      </List>
    </Container>
  </Wrapper >
)

export default DashboardEdit
