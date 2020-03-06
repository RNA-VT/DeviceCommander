import * as React from 'react'
import { Component, MouseEvent } from 'react'
import Solenoid from '../../utils/Solenoid'
import SolenoidButton from './SolenoidButton'
import DeviceManagement from '../../containers/DeviceManagementContainer'
import SolenoidFactory from '../../utils/factories/SolenoidFactory'
import styled from 'styled-components'
import { MenuItem, InputLabel, Select } from '@material-ui/core'

const TitleRow = styled.div`
  // background-color: gray;
`

type PBCProps = {
  deviceManager: DeviceManagement
}

type PBCState = {
  buttonConfigs: Array<any>,
  deviceManager: DeviceManagement
}

class PositionedButtonContainer extends Component<PBCProps, PBCState> {
  constructor(props: PBCProps) {
    super(props)

    console.log('props', props)

    this.state = {
      buttonConfigs: [],
      deviceManager: props.deviceManager
    }
  }

  handleAddButton(e: React.ChangeEvent<{
    name?: string | undefined;
    value: unknown;
  }>, target: any) {
    if (target) {
      console.log('handleAddButton', target.props.solenoid)
    }
  }


  render() {
    console.log('state', this.state)
    let solenoidComponents: Array<any> = []
    let solenoidListItems: Array<any> = []
    let sf = new SolenoidFactory

    if (this.state.deviceManager) {
      const solenoids = sf.makeSolenoidsFromManyMcs(this.state.deviceManager.getMicrocontrollers())

      let tmpXPos = 0
      let tmpYPos = 0
      solenoidComponents = solenoids.map(solenoid => {
        tmpXPos += 20
        tmpYPos += 100
        return (
          <SolenoidButton
            key={solenoid.uid}
            solenoid={solenoid}
            xPos={tmpXPos}
            yPos={tmpYPos} />
        )
      })

      solenoidListItems = solenoids.map(solenoid => {
        return (
          <MenuItem key={solenoid.uid}
          // solenoid={solenoid}
          >{solenoid.name}
          </MenuItem>
        )
      })

    }

    console.log('solenoids', solenoidComponents)

    return (
      <div>
        <TitleRow>
          <h1>PositionedButtonContainer</h1>
          <InputLabel id="add-btn">Add a Button</InputLabel>
          <Select labelId="add-btn" id="select" onChange={this.handleAddButton}>
            {solenoidListItems}
          </Select>
        </TitleRow>
        <div>
          {solenoidComponents}
        </div>
      </div>
    )
  }
}

export default PositionedButtonContainer