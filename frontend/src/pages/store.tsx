import { atom, selector, selectorFamily } from 'recoil';
import GetDevicesMethod from '../api/method/GetDevices';
import Device from '../api/device/Device';

interface DeviceListState {
  devices: Device[]
  loadCount: number
}

export const DeviceListAtom = atom<DeviceListState>({
  key: 'deviceListAtom',
  default: {
    devices: [],
    loadCount: 0,
  },
});

export const DevicesState = selectorFamily<Device[], number>({
  key: 'devices',
  get: (_) => async () => {
    const apiMethod = new GetDevicesMethod('http://localhost:8001');
    const devices = await apiMethod.do();

    return devices;
  },
});

export const currentDeviceListQuery = selector({
  key: 'currentDeviceListQuery',
  get: ({ get }) => {
    const devices = get(DevicesState(get(DeviceListAtom).loadCount));
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
