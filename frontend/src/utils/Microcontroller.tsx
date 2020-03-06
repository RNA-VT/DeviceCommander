import Solenoid from "./Solenoid"

class Microcontroller {
    id: string | undefined
    host: string | undefined
    port: string | undefined
    description: string | undefined
    solenoids: Array<Solenoid> | undefined
}

export default Microcontroller