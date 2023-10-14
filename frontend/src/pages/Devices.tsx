import React, { useEffect, useState } from 'react';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import ThumbDownIcon from '@mui/icons-material/ThumbDown';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import CachedIcon from '@mui/icons-material/Cached';
import DeleteIcon from '@mui/icons-material/Delete';
import Stack from '@mui/material/Stack';
import { Button } from '@mui/material';
import { useRecoilValue, useSetRecoilState, useRecoilRefresher_UNSTABLE } from 'recoil';
import Device from '../api/device/Device';
import Dashboard from '../layouts/dashboard/Dashboard';
import { PageState, DevicesState } from './store';
import DeleteDevicesMethod from '../api/method/DeleteDevice';
import RunARPScanMethod from '../api/cluster/RunARPScan';
import AddDeviceModal from '../modals/AddDevice';

export default function Devices() {
  const setPageState = useSetRecoilState(PageState);
  const devices = useRecoilValue(DevicesState);
  const refreshDeviceList = useRecoilRefresher_UNSTABLE(DevicesState);
  const [addDeviceModalOpen, setAddDeviceModalOpen] = useState(false);

  const deleteDevice = async (id: string) => {
    console.log(`Deleting device ${id}`);

    const deleteMethod = new DeleteDevicesMethod('http://localhost:8001', id);
    const resp = await deleteMethod.do();
    console.log(resp);

    refreshDeviceList();
  };

  const handleDeviceScanClick = () => {
    console.log('Device Scan');
    const scanMethod = new RunARPScanMethod('http://localhost:8001');
    const resp = scanMethod.do();
    console.log(resp);
  };

  const handleAddDeviceClick = () => {
    console.log('Add Device');
    setAddDeviceModalOpen(true);
  };

  const handleAddDeviceModalClose = () => {
    setAddDeviceModalOpen(false);
    refreshDeviceList();
  };

  useEffect(() => {
    console.log('Setting page state');
    setPageState({
      title: 'Devices',
      index: 'devices',
    });
  });

  return (
    <Dashboard>
      <Container maxWidth="xl" sx={{ mt: 4, mb: 4 }}>
        <Stack direction="row" justifyContent="flex-end" spacing={1}>
          <Button variant="contained" onClick={handleDeviceScanClick}>Run Device Scan</Button>
          <Button variant="contained" onClick={handleAddDeviceClick}>Add Device</Button>
          <Button variant="contained" onClick={refreshDeviceList}><CachedIcon /></Button>
          <AddDeviceModal open={addDeviceModalOpen} handleClose={handleAddDeviceModalClose} />
        </Stack>
      </Container>
      <Container maxWidth="xl" sx={{ mt: 4, mb: 4 }}>
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">MAC Address</TableCell>
                <TableCell align="right">IP</TableCell>
                <TableCell align="center">Failures</TableCell>
                <TableCell align="center">Active</TableCell>
                <TableCell align="center">Delete</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {devices.map((device: Device) => (
                <TableRow
                  key={device.ID}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell component="th" scope="row">
                    {device.Name}
                  </TableCell>
                  <TableCell align="right">{device.MAC}</TableCell>
                  <TableCell align="right">
                    {device.Host}
                    :
                    {device.Port}
                  </TableCell>
                  <TableCell align="center">{device.Failures}</TableCell>
                  <TableCell align="center">{device.Active ? <ThumbUpIcon /> : <ThumbDownIcon />}</TableCell>
                  <TableCell align="center">
                    <Button onClick={(_: React.MouseEvent<HTMLElement>) => {
                      deleteDevice(device.ID);
                    }}
                    >
                      <DeleteIcon />
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </Dashboard>
  );
}
