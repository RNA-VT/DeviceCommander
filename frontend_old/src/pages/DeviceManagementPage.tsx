import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import DeviceCard from '../components/device_management/DeviceCard'
import Device from '../utils/Device'

const DeviceManagementPage = () => {


  return (
    <Wrapper>
      <Container maxWidth="md">
        <h1>Device Management Page</h1>
        <List>
          <Subscribe to={[DeviceManagement]}>

            {deviceManager => {
              const devs: Array<Device> = deviceManager.getDevices()

              const handleReload = () => {
                return deviceManager.loadData()
              }

              return devs.map<React.ReactNode>((dev) => (
                <ListItem key={dev.id}>
                  <DeviceCard device={dev} reload={handleReload} />
                </ListItem>
              ))
            }}
          </Subscribe>
        </List>
      </Container>
    </Wrapper >)
}

export default DeviceManagementPage
