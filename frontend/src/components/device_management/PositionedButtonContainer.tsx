import * as React from 'react'
import { Component } from 'react'
import { Subscribe } from 'unstated-typescript'
import SolenoidButton from './SolenoidButton'
import DeviceManagement from '../../containers/DeviceManagementContainer'
import ControlPanelContainer from '../../containers/ControlPanelContainer'
import SolenoidFactory from '../../utils/factories/SolenoidFactory'
import styled from 'styled-components'
import {
  MenuItem,
  InputLabel,
  Select,
  Grid,
  Button,
  FormControl
} from '@material-ui/core'

const TitleRow = styled.div`
  // background-color: gray;
`

const FormControlStyled = styled(FormControl)`
  min-width: 140px;
  width: 100%;
`

const SubmitButton = styled(Button)`
  height: 100%;
`

type PBCProps = {
  deviceManager: DeviceManagement
}

type PBCState = {
  buttonConfigs: Array<any>,
  deviceManager: DeviceManagement,
  componentSelect: any,
  controlTypeSelect: any
}

class PositionedButtonContainer extends Component<PBCProps, PBCState> {
  constructor(props: PBCProps) {
    super(props)

    this.state = {
      buttonConfigs: [],
      deviceManager: props.deviceManager,
      componentSelect: "",
      controlTypeSelect: ""
    }

    this.handleChangeComponentButton = this.handleChangeComponentButton.bind(this)
    this.handleChangeControlType = this.handleChangeControlType.bind(this)
    this.handleAddControl = this.handleAddControl.bind(this)
  }

  handleChangeComponentButton(e: React.ChangeEvent<{
    name?: string | undefined;
    value: unknown;
  }>, target: any) {
    if (target) {
      console.log('handleAddButton', e)
      this.setState({
        componentSelect: e.target.value
      })
    }

  }

  handleChangeControlType(e: React.ChangeEvent<{
    name?: string | undefined;
    value: unknown;
  }>, target: any) {
    if (target) {
      console.log('handleAddButton', e)
      this.setState({
        controlTypeSelect: e.target.value
      })
    }
  }

  handleAddControl() {

  }

  handleBtnMove(uid: string, xPos: number, yPos: number) {

  }

  render() {
    let solenoidListItems: Array<any> = []
    let tmpSolenoidBtnConfigs: Array<any> = []

    if (this.state.deviceManager) {
      const solenoids = this.state.deviceManager.getSolenoids()

      let tmpXPos = 0
      let tmpYPos = 0


      solenoidListItems = solenoids.map(solenoid => {
        return (
          <MenuItem key={solenoid.uid} value={solenoid.uid}>{solenoid.name}</MenuItem>
        )
      })

    }

    console.log(solenoidListItems)

    return (
      <div>
        <div>
          <TitleRow>
            <h1>PositionedButtonContainer</h1>


          </TitleRow>
          <Grid container spacing={1}>
            <Grid item xs={3}>
              <FormControlStyled>
                <InputLabel disableAnimation={true}
                  id="add-btn-label">Add a Button</InputLabel>
                <Select labelId="add-btn-label"
                  value={this.state.componentSelect}
                  onChange={this.handleChangeComponentButton}>
                  {solenoidListItems}
                </Select>
              </FormControlStyled>
            </Grid>
            <Grid item xs={3}>
              <FormControlStyled>
                <InputLabel disableAnimation={true}
                  id="control-type-label">Control Type</InputLabel>
                <Select labelId="control-type-label"
                  value={this.state.controlTypeSelect}
                  onChange={this.handleChangeControlType}>
                  <MenuItem value={"button"}>Button</MenuItem>
                  <MenuItem value={"switch"}>Switch</MenuItem>
                </Select>
              </FormControlStyled>
            </Grid>
            <Grid item xs={3}>
              <SubmitButton type="submit"
                variant="outlined"
                onClick={this.handleAddControl}>Add Control</SubmitButton>
            </Grid>
          </Grid>

        </div>

        <div>
          <Subscribe to={[ControlPanelContainer]}>
            {controlPanelContainer => {
              return controlPanelContainer.getConfigs().map<React.ReactNode>((config) => (
                <SolenoidButton
                  key={config.componentUID}
                  xPos={config.xPos}
                  yPos={config.yPos}
                  setPosition={this.handleAddControl} />
              ))
            }}
          </Subscribe>
        </div>
      </div>
    )
  }
}

export default PositionedButtonContainer