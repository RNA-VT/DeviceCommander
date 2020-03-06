import Solenoid from "../../utils/Solenoid"

class SolenoidFactory {
  constructor() {

  }

  makeSolenoidsFromManyMcs(mcs: Array<any>) {
    let allSolenoids: Solenoid[] = []

    mcs.forEach(mc => {
      allSolenoids = allSolenoids.concat(this.makeSolenoidsFromMc(mc))
    })

    return allSolenoids
  }

  makeSolenoidsFromMc(mc: any) {
    return mc.Solenoids.map((solenoid: any) => {
      return new Solenoid(solenoid, mc)
    })
  }
}

export default SolenoidFactory