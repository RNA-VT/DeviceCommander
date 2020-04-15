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

  getConfigs(): Array<ControlConfig> {
    return this.state.controlConfigs
  }
}

export default ControlPanelContainer