import { Container } from 'unstated-typescript'

type ControlConfig = {
  componentUID: string,
  controlType: string,
  inputType: string,
  xPos: number,
  yPos: number
}

type ControlPanelState = {
  controlConfigs: Array<ControlConfig>
}

class ControlPanelContainer extends Container<ControlPanelState> {
  constructor() {
    super()

    this.state = {
      controlConfigs: []
    }

    this.setControlPosition = this.setControlPosition.bind(this)
  }

  addButton(config: ControlConfig) {
    this.setState({
      controlConfigs: [...this.state.controlConfigs, config]
    })

  }

  setConfigs(data: Array<ControlConfig>) {
    this.setState({
      controlConfigs: data
    })
  }

  setControlPosition(uid: string, xPos: number, yPos: number) {
    console.log('setControlPosition', this.state)
    this.setState({
      controlConfigs: this.state.controlConfigs.map((control) => {
        if (control.componentUID == uid) {
          control.xPos = xPos
          control.yPos = yPos
        }
        return control
      })
    })

  }

  getConfigs(): Array<ControlConfig> {
    return this.state.controlConfigs
  }
}

export default ControlPanelContainer