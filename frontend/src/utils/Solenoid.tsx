import API from "../utils/ApiWrapper"

class Solenoid {
  uid: string
  name: string
  mcAddress: string
  type: string
  apiRoot: API
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

    this.apiRoot = new API(this.mcAddress)

    this.open = this.open.bind(this)
    this.close = this.close.bind(this)
  }

  open() {
    this.apiRoot.openSolenoid(this.uid)
  }

  close() {
    this.apiRoot.closeSolenoid(this.uid)
  }

  getContructorData() {
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