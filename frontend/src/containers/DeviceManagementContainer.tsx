import { Container } from 'unstated-typescript'
import API from '../utils/ApiWrapper'
import Device from '../utils/Device'
import DeviceFactory from "../utils/factories/DeviceFactory"
import Solenoid from '../utils/Solenoid'

type DeviceManagementState = {
  isLoaded: boolean,
  devices: Array<Device>,
  master?: Device,
  clusterName: String
}

class DeviceManagement extends Container<DeviceManagementState> {
  constructor() {
    super()
    this.state = {
      isLoaded: false,
      devices: [],
      clusterName: ''
    }

    this.getDevices = this.getDevices.bind(this)
    this.getData = this.getData.bind(this)
    this.loadData()
  }

  async getData() {
    const api = new API("")
    return api.getClusterInfo()
  }

  async loadData() {
    return this.getData().then((data) => {
      console.log('load data', data);
      let devFactory = new DeviceFactory()

      this.setState({
        devices: devFactory.makeManyDevices(data.Devices),
        clusterName: data.Name
      })
    })
  }

  getDevices(): Array<Device> {
    if (this.state.devices) {
      return [...this.state.devices]
    }

    return []
  }

  getSolenoids(): Array<Solenoid> {
    let allSolenoids: Solenoid[] = []
    this.getDevices().forEach(dev => {
      allSolenoids = allSolenoids.concat(dev.solenoids)
    })
    return allSolenoids
  }


}

export default DeviceManagement
