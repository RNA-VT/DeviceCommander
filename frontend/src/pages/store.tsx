import { atom, selector } from 'recoil';
import GetDevicesMethod from '../api/method/GetDevices';
import Device from '../api/device/Device';

interface PageData {
  title: string
  index: string
}

// export const DevicesState = atom<Device[]>({
//   key: "devices",
//   default: [],
// });

export const DevicesState = selector<Device[]>({
  key: 'devices',
  get: async () => {
    // const response = await myDBQuery({
    //   userID: get(currentUserIDState),
    // });
    const apiMethod = new GetDevicesMethod('http://localhost:8001');
    const devices = await apiMethod.do();

    return devices;
  },
});

export const PageState = atom<PageData>({
  key: 'pageState',
  default: {
    title: 'Home',
    index: 'home',
  },
});
