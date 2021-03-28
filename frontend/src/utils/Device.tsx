import Solenoid from "./Solenoid"
import SolenoidFactory from "./factories/SolenoidFactory"
import ApiWrapper from "./ApiWrapper"

class Device {
  id: string;
  host: string | undefined;
  name: string | undefined;
  port: string | undefined;
  description: string | undefined;
  solenoids: Array<Solenoid>;

  constructor(data: any) {
    this.id = data.ID
    this.name = data.Name
    this.description = data.Description
    this.host = data.Host
    this.port = data.Port

    if (data.Solenoids) {
      const sf = new SolenoidFactory()
      this.solenoids = sf.makeSolenoidsFromDevice(data)
    } else {
      this.solenoids = []
    }
  }

  myNetworkAddress(): string {
    return this.host + ':' + this.port
  }

  async edit(newData: any) {
    const api = new ApiWrapper(this.myNetworkAddress())
    return api.editDevice(this.id, newData)
  }
}

export default Device