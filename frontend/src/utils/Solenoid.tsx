import ApiWrapper from "../utils/ApiWrapper"

class Solenoid {
  uid: string
  name: string
  mcAddress: string
  type: string
  enabled: boolean
  headerPin: number

  constructor(solenoid: any, mc: any) {
    this.uid = solenoid.UID ? solenoid.UID : ''
    this.name = solenoid.Name ? solenoid.Name : ''
    this.type = solenoid.Type ? solenoid.Type : ''
    this.enabled = solenoid.Enabled
    this.headerPin = solenoid.HeaderPin

    if (solenoid.mcAddress) {
      this.mcAddress = solenoid.mcAddress
    } else {
      this.mcAddress = (mc.Host && mc.Port) ? (mc.Host + ':' + mc.Port) : ''
    }

    this.open = this.open.bind(this)
    this.close = this.close.bind(this)
  }

  open() {
    const api = new ApiWrapper(this.mcAddress)
    api.openSolenoid(this.uid)
  }

  close() {
    const api = new ApiWrapper(this.mcAddress)
    api.closeSolenoid(this.uid)
  }

  async edit(newData: any) {
    const api = new ApiWrapper(this.mcAddress)
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
      mcAddress: this.mcAddress
    }
  }
}

export default Solenoid