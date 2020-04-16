import * as React from 'react'
import { Component } from 'react'
import { Subscribe } from 'unstated-typescript'
import ControlButton from '../control_panel/ControlButton'
import ControlSwitch from '../control_panel/ControlSwitch'
import DeviceManagement from '../../containers/DeviceManagementContainer'
import ControlPanelContainer, { ControlConfig } from '../../containers/ControlPanelContainer'
import ControlPanelConfigEdit from '../control_panel/ControlPanelConfigEdit'
import SolenoidFactory from '../../utils/factories/SolenoidFactory'
import styled from 'styled-components'
import {
  MenuItem,
  InputLabel,
  Select,
  Grid,
  Button,
  FormControl,
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
  width: 100%;
`

const ControlPanelGrid = styled.div`
  height: 1000px;
  width: 1000px;
  overflow: hidden;
  border: black 2px solid;
  margin: 10px auto;
`

type PBCProps = {
  deviceManager: DeviceManagement,
  controlPanelManager: ControlPanelContainer
}

type PBCState = {
  buttonConfigs: Array<any>,
  deviceManager: DeviceManagement,
  controlPanelManager: ControlPanelContainer,
  componentSelect: any,
  controlTypeSelect: any,
  jsonDialogOpen: boolean
}

class PositionedButtonContainer extends Component<PBCProps, PBCState> {
  constructor(props: PBCProps) {
    super(props)

    this.state = {
      buttonConfigs: [],
      deviceManager: props.deviceManager,
      controlPanelManager: props.controlPanelManager,
      componentSelect: "",
      controlTypeSelect: "",
      jsonDialogOpen: false
    }

    this.handleChangeComponentButton = this.handleChangeComponentButton.bind(this)
    this.handleChangeControlType = this.handleChangeControlType.bind(this)
    this.handleAddControl = this.handleAddControl.bind(this)
    this.handleExportConfig = this.handleExportConfig.bind(this)
    this.handleCloseJsonDialog = this.handleCloseJsonDialog.bind(this)
    this.handleExportConfig = this.handleExportConfig.bind(this)
  }

  handleChangeComponentButton(e: React.ChangeEvent<{
    name?: string | undefined;
    value: unknown;
  }>, target: any) {
    if (target) {
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
      this.setState({
        controlTypeSelect: e.target.value
      })
    }
  }

  handleAddControl() {
    const newBtn = {
      componentUID: this.state.componentSelect,
      controlType: 'something',
      inputType: this.state.controlTypeSelect,
      xPos: 0,
      yPos: 0
    }
    this.state.controlPanelManager.addButton(newBtn)
  }

  handleCloseJsonDialog() {
    this.setState({
      jsonDialogOpen: false
    })
  }

  handleExportConfig() {
    console.log('Control Panel Config', this.state.controlPanelManager.getConfigs())
    this.setState({
      jsonDialogOpen: true
    })
  }

  render() {
    let solenoidListItems: Array<any> = []
    if (this.state.deviceManager) {
      const solenoids = this.state.deviceManager.getSolenoids()
      solenoidListItems = solenoids.map(solenoid => {
        return (
          <MenuItem key={solenoid.uid} value={solenoid.uid}>{solenoid.name}</MenuItem>
        )
      })
    }

    let controlList: Array<any> = []
    if (this.state.controlPanelManager) {
      controlList = this.state.controlPanelManager.getConfigs().map(control => {
        if (control.inputType == 'switch') {
          return (
            <ControlSwitch
              key={control.componentUID}
              componentUID={control.componentUID}
              xPos={control.xPos}
              yPos={control.yPos}
              label={control.componentUID}
              setPosition={this.state.controlPanelManager.setControlPosition} />
          )
        } else {
          return (
            <ControlButton
              key={control.componentUID}
              componentUID={control.componentUID}
              xPos={control.xPos}
              yPos={control.yPos}
              label={control.componentUID}
              setPosition={this.state.controlPanelManager.setControlPosition} />
          )

        }
      })
    }

    return (
      <div>
        <div>
          <TitleRow>
            <h1>PositionedButtonContainer</h1>
          </TitleRow>
          <Grid container spacing={1}>
            <Grid item xs={4}>
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
            <Grid item xs={4}>
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
            <Grid item xs={2}>
              <SubmitButton type="submit"
                variant="outlined"
                onClick={this.handleAddControl}>Add Control</SubmitButton>
            </Grid>
          </Grid>
        </div>
        <ControlPanelGrid>
          {controlList}
        </ControlPanelGrid>
        <ControlPanelConfigEdit
          jsonConfig={JSON.stringify(this.state.controlPanelManager.getConfigs(), null, 2)}
          setControlPanelConfig={this.state.controlPanelManager.setConfigs}
        />
      </div>
    )
  }
}

export default PositionedButtonContainer