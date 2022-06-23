import Solenoid from "../../utils/Solenoid"

class SolenoidFactory {

  makeSolenoidsFromManyDevs(devs: Array<any>) {
    let allSolenoids: Solenoid[] = []
    devs.forEach(dev => {
      allSolenoids = allSolenoids.concat(this.makeSolenoidsFromDev(dev))
    })

    return allSolenoids
  }

  makeSolenoidsFromDev(dev: any) {
    if (dev.solenoids) {
      return dev.solenoids.map((solenoid: any) => {
        return new Solenoid(solenoid, dev)
      })
    }

    if (dev.Solenoids) {
      return dev.Solenoids.map((solenoid: any) => {
        return new Solenoid(solenoid, dev)
      })
    }

    return []

  }
}

export default SolenoidFactory