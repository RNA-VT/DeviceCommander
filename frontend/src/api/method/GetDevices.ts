import { APIMethod } from './Method'
import { Device } from '../device/Device'
import axios from 'axios'


export class GetDevicesMethod implements APIMethod {
    baseApiURL:string
    constructor(baseApiURL:string) {
        this.baseApiURL = baseApiURL
     }

    async do(): Promise<Device[]> {
        const { data, status } = await axios.get<Device[]>(this.baseApiURL + "/v1/device", {
            headers: {
                Accept: 'application/json',
            },
        })

        return data
    }
}