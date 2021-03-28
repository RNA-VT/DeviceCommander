import Device from "../Device"

class DeviceFactory {

    makeManyDevices(devs: Array<any>) {
        let allMicrocontrollers: Device[] = []

        devs.forEach(dev => {

            allMicrocontrollers = allMicrocontrollers.concat(new Device(dev))
        })

        return allMicrocontrollers
    }
}

export default DeviceFactory