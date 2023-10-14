import axios from 'axios';
import { APIMethod } from '../Method';

export interface ARPScanResponse {
  Message: string;
}

export default class RunARPScanMethod implements APIMethod {
  baseApiURL: string;

  constructor(baseApiURL: string) {
    this.baseApiURL = baseApiURL;
  }

  async do(): Promise<ARPScanResponse> {
    const { data } = await axios.get<ARPScanResponse>(
      `${this.baseApiURL}/v1/ops/run-arp-scan`,
      {
        headers: {
          Accept: 'application/json',
        },
      },
    );

    return data;
  }
}
