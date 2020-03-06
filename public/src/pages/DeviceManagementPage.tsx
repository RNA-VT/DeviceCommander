import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { Link } from 'react-router-dom'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'
import DeviceCard from '../components/device_management/DeviceCard'


const DeviceManagementPage = () => {
  return (
    <Wrapper>
      <PageHeader />
      <Container maxWidth="md">
        <h1>Device Management Page</h1>
        <List>
          <Subscribe to={[DeviceManagement]}>

            {deviceManager =>
              deviceManager.getMicrocontrollers().map((mc: any) => (
                <>
                  <ListItem key={mc.ID}>
                    <DeviceCard microcontroller={mc} />
                  </ListItem>
                </>
              ))
            }
          </Subscribe>
        </List>
      </Container>
    </Wrapper >)
}

export default DeviceManagementPage
