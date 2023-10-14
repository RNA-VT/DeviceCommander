import axios from 'axios';
import { APIMethod } from '../Method';
import Device from '../device/Device';

export default class GetDevicesMethod implements APIMethod {
  baseApiURL:string;

  constructor(baseApiURL:string) {
    this.baseApiURL = baseApiURL;
  }

  async do(): Promise<Device[]> {
    const { data } = await axios.get<Device[]>(`${this.baseApiURL}/v1/device`, {
      headers: {
        Accept: 'application/json',
      },
    });

    return data;
  }
}
