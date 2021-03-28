import Solenoid from "../../utils/Solenoid"

class SolenoidFactory {

  makeSolenoidsFromManyMcs(mcs: Array<any>) {
    let allSolenoids: Solenoid[] = []
    mcs.forEach(mc => {
      allSolenoids = allSolenoids.concat(this.makeSolenoidsFromMc(mc))
    })

    return allSolenoids
  }

  makeSolenoidsFromMc(mc: any) {
    if (mc.solenoids) {
      return mc.solenoids.map((solenoid: any) => {
        return new Solenoid(solenoid, mc)
      })
    }
    
    return []
  }
}

export default SolenoidFactory