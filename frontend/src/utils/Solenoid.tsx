import ApiWrapper from "../utils/ApiWrapper"

class Solenoid {
  uid: string
  name: string
  devAddress: string
  type: string
  enabled: boolean
  headerPin: number

  constructor(solenoid: any, dev: any) {
    this.uid = solenoid.UID ? solenoid.UID : ''
    this.name = solenoid.Name ? solenoid.Name : ''
    this.type = solenoid.Type ? solenoid.Type : ''
    this.enabled = solenoid.Enabled
    this.headerPin = solenoid.HeaderPin

    if (solenoid.devAddress) {
      this.devAddress = solenoid.devAddress
    } else {
      this.devAddress = (dev.Host && dev.Port) ? (dev.Host + ':' + dev.Port) : ''
    }

    this.open = this.open.bind(this)
    this.close = this.close.bind(this)
  }

  open() {
    const api = new ApiWrapper(this.devAddress)
    api.openSolenoid(this.uid)
  }

  close() {
    const api = new ApiWrapper(this.devAddress)
    api.closeSolenoid(this.uid)
  }

  async edit(newData: any) {
    const api = new ApiWrapper(this.devAddress)
    return api.editComponent(this.uid, newData)
  }

  // Get solenoid data into the shape it needs to be to construct
  getConfig() {
    return {
      UID: this.uid,
      Name: this.name,
      Type: this.type,
      Enabled: this.enabled,
      HeaderPin: this.headerPin,
      devAddress: this.devAddress
    }
  }
}

export default Solenoid