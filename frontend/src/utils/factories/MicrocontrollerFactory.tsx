import Microcontroller from "../../utils/Microcontroller"

class MicrocontrollerFactory {

    makeManyMcs(mcs: Array<any>) {
        let allMicrocontrollers: Microcontroller[] = []

        mcs.forEach(mc => {

            allMicrocontrollers = allMicrocontrollers.concat(new Microcontroller(mc))
        })

        return allMicrocontrollers
    }
}

export default MicrocontrollerFactory