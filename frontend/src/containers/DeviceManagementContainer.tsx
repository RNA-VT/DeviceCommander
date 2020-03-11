import { Container } from 'unstated-typescript'
import API from '../utils/ApiWrapper'
import Microcontroller from '../utils/Microcontroller'
import MicrocontrollerFactory from "../utils/factories/MicrocontrollerFactory"
import Solenoid from '../utils/Solenoid'

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
    this.loadData()
  }

  async getData() {
    const api = new API("http://localhost:8000")
    return api.getClusterInfo()
  }

  async loadData() {
    console.log('LOAD DATA')
    return this.getData().then((data) => {
      console.log("SETTING STATE");
      let mcFactory = new MicrocontrollerFactory()
      console.log(data);

      this.setState({
        slaveMicrocontrollers: mcFactory.makeManyMcs(data.SlaveMicrocontrollers),
        master: new Microcontroller(data.Master),
        clusterName: data.Name
      })
    })
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

  getSolenoids(): Array<Solenoid> {
    let allSolenoids: Solenoid[] = []
    this.getMicrocontrollers().forEach(mc => {
      allSolenoids = allSolenoids.concat(mc.solenoids)
    })
    return allSolenoids
  }


}

export default DeviceManagement
