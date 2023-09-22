import axios from 'axios';
import { APIMethod } from './Method';
import Device from '../device/Device';

export default class DeleteDevicesMethod implements APIMethod {
  baseApiURL: string;

  targetID: string;

  constructor(baseApiURL: string, id: string) {
    this.baseApiURL = baseApiURL;
    this.targetID = id;
  }

  async do(): Promise<Device[]> {
    const { data } = await axios.delete<Device[]>(`${this.baseApiURL}/v1/device/${this.targetID}`, {
      headers: {
        Accept: 'application/json',
      },
    });

    return data;
  }
}
