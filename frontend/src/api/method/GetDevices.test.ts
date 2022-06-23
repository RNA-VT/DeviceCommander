import {GetDevicesMethod} from './GetDevices'

test('test getDevices', async () => {
    let method = new GetDevicesMethod("http://localhost:8001");

    const response = await method.do()
    expect(response.length).toBeGreaterThan(0)
})