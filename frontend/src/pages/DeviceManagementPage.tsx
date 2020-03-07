import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import DeviceCard from '../components/device_management/DeviceCard'

const DeviceManagementPage = () => {
  return (
    <Wrapper>
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
