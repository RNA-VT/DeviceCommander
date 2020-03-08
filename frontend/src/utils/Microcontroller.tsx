import Solenoid from "./Solenoid"
import SolenoidFactory from "./factories/SolenoidFactory"

class Microcontroller {
  id: string | undefined;
  host: string | undefined;
  name: string | undefined;
  port: string | undefined;
  description: string | undefined;
  solenoids: Array<Solenoid>;


  constructor(data: any) {
    console.log('MC Constructor', data)
    this.id = data.ID
    this.name = data.Name
    this.description = data.Description
    this.host = data.Host
    this.port = data.Port


    if (data.Solenoids) {
      const sf = new SolenoidFactory()
      this.solenoids = sf.makeSolenoidsFromMc(data)
    } else {
      this.solenoids = []
    }

  }
}

export default Microcontroller