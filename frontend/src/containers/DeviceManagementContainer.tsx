import { Container } from 'unstated-typescript'
import API from '../utils/ApiWrapper'
import Microcontroller from '../utils/Microcontroller'
import MicrocontrollerFactory from "../utils/factories/MicrocontrollerFactory"

type DeviceManagementState = {
  isLoaded: boolean,
  slaveMicrocontrollers: Array<Microcontroller>,
  master?: Microcontroller,
  clusterName: String
}

class DeviceManagement extends Container<DeviceManagementState> {
  constructor() {
    super()
    this.state = {
      isLoaded: false,
      master: undefined,
      slaveMicrocontrollers: [],
      clusterName: ''
    }

    this.getMicrocontrollers = this.getMicrocontrollers.bind(this)
    this.getData = this.getData.bind(this)
    this.getData().then((data) => {
      let mcFactory = new MicrocontrollerFactory()
      console.log(data);

      this.setState({
        slaveMicrocontrollers: mcFactory.makeManyMcs(data.SlaveMicrocontrollers),
        master: new Microcontroller(data.Master),
        clusterName: data.Name
      })
    })
  }

  async getData() {
    const api = new API("http://localhost:8000")
    return api.getClusterInfo()
  }

  getMicrocontrollers(): Array<Microcontroller> {
    if (this.state.slaveMicrocontrollers) {
      return [...this.state.slaveMicrocontrollers]
    }

    if (this.state.master) {
      return [this.state.master]
    }

    return []
  }


}

export default DeviceManagement
