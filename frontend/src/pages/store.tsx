import { atom, selector } from 'recoil';
import GetDevicesMethod from '../api/method/GetDevices';
import Device from '../api/device/Device';

export const DevicesState = selector<Device[]>({
  key: 'devices',
  get: async () => {
    const apiMethod = new GetDevicesMethod('http://localhost:8001');
    const devices = await apiMethod.do();

    return devices;
  },
});

interface PageData {
  title: string
  index: string
}

export const PageState = atom<PageData>({
  key: 'pageState',
  default: {
    title: 'Home',
    index: 'home',
  },
});
