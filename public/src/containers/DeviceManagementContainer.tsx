import { Container } from 'unstated-typescript'
import API from '../utils/apiWrapper'

type DeviceManagementState = {
  isLoaded: boolean,
  slaveMicrocontrollers: Array<object>,
  master: Array<object>,
  clusterName: String
}

class DeviceManagement extends Container<DeviceManagementState> {
  constructor() {
    super()
    this.state = {
      isLoaded: false,
      master: [],
      slaveMicrocontrollers: [],
      clusterName: ''
    }

    this.getMicrocontrollers = this.getMicrocontrollers.bind(this)
    this.getData = this.getData.bind(this)
    this.getData().then((data) => {
      console.log(data);
      this.setState({
        slaveMicrocontrollers: data.SlaveMicrocontrollers,
        master: data.Master,
        clusterName: data.Name
      })
    })
  }

  async getData() {

    const api = new API("")
    return api.getClusterInfo()
  }

  getMicrocontrollers() {
    if (this.state.slaveMicrocontrollers) {
      return [...this.state.slaveMicrocontrollers]
    }

    return [this.state.master]
  }


}

export default DeviceManagement
