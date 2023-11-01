import axios from 'axios';
import { APIMethod } from '../Method';
import Device from '../device/Device';

export interface NewDeviceData {
  host: string;
  port: number;
  name: string;
  description: string;
}

export default class AddDeviceMethod implements APIMethod {
  baseApiURL: string;

  newDevice: NewDeviceData;

  constructor(baseApiURL: string, newDevice: NewDeviceData) {
    this.baseApiURL = baseApiURL;
    this.newDevice = newDevice;
  }

  async do(): Promise<Device> {
    const { data } = await axios.post<Device>(
      `${this.baseApiURL}/v1/device`,
      {
        Name: this.newDevice.name,
        Host: this.newDevice.host,
        Port: this.newDevice.port,
      },
      {
        headers: {
          Accept: 'application/json',
        },
      },
    );

    return data;
  }
}
