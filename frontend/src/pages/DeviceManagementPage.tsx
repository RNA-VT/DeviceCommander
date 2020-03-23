import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import DeviceCard from '../components/device_management/DeviceCard'
import Microcontroller from '../utils/Microcontroller'

const DeviceManagementPage = () => {


  return (
    <Wrapper>
      <Container maxWidth="md">
        <h1>Device Management Page</h1>
        <List>
          <Subscribe to={[DeviceManagement]}>

            {deviceManager => {
              const mcs: Array<Microcontroller> = deviceManager.getMicrocontrollers()

              const handleReload = () => {
                return deviceManager.loadData()
              }

              return mcs.map<React.ReactNode>((mc) => (
                <ListItem key={mc.id}>
                  <DeviceCard microcontroller={mc} reload={handleReload} />
                </ListItem>
              ))
            }}
          </Subscribe>
        </List>
      </Container>
    </Wrapper >)
}

export default DeviceManagementPage
