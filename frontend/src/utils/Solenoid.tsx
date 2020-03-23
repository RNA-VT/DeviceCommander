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
    this.mcAddress = (mc.Host && mc.Port) ? (mc.Host + ':' + mc.Port) : ''
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
}

export default Solenoid