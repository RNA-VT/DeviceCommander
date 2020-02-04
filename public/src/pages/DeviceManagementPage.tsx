import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import { Link } from 'react-router-dom'
import { List, ListItem, Container } from '@material-ui/core'
import DeviceManagement from '../containers/DeviceManagementContainer'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'

import SolenoidFactory from '../utils/factories/SolenoidFactory'
import Solenoid from '../utils/Solenoid'

import SolenoidButton from '../components/device_management/SolenoidButton'
import DeviceCard from '../components/device_management/DeviceCard'


import Draggable, { DraggableCore } from "react-draggable"


const DeviceManagementPage = () => {
  return (
    <Wrapper>
      <PageHeader />
      <Container maxWidth="md">
        <h1>Device Management Page</h1>
        <List>
          <Subscribe to={[DeviceManagement]}>

            {deviceManager =>
              deviceManager.getMicrocontrollers().map((mc) => (
                <ListItem key={mc.ID}>
                  <DeviceCard>
                    {mc.Host}:{mc.Port}

                    {
                      mc.Solenoids.map((solenoid) => {
                        return (
                          <div key={solenoid.UID}>
                            {solenoid.Name}
                          </div>
                        )
                      })
                    }
                  </DeviceCard>

                </ListItem>
              ))
            }

            {/* {deviceManager => {
              console.log(deviceManager.getMicrocontrollers())
              let sf = new SolenoidFactory
              const solenoids = sf.makeSolenoidsFromManyMcs(deviceManager.getMicrocontrollers())
              console.log('solenoids', solenoids)

              const solenoidComponents = solenoids.map(solenoid => {
                return (
                  <SolenoidButton solenoid={solenoid} />
                )
              })

              return (
                <div>
                  {solenoidComponents}
                </div>
              )
            }} */}
          </Subscribe>
        </List>
      </Container>
    </Wrapper >)
}

export default DeviceManagementPage
