import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { Link } from 'react-router-dom'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'
import DeviceCard from '../components/device_management/DeviceCard'
import SolenoidFactory from '../utils/factories/SolenoidFactory'
import SolenoidButton from '../components/device_management/SolenoidButton'

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
