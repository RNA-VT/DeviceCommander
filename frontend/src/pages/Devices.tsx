import React, { useEffect } from 'react';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import CheckIcon from '@mui/icons-material/Check';
import ClearIcon from '@mui/icons-material/Clear';
import { useRecoilValue, useSetRecoilState } from 'recoil';
import Device from '../api/device/Device';
import Dashboard from '../layouts/dashboard/Dashboard';
import { PageState, DevicesState } from './store';

export default function Devices() {
  const devices = useRecoilValue(DevicesState);
  const setPageState = useSetRecoilState(PageState);

  useEffect(() => {
    setPageState({
      title: 'Devices',
      index: 'devices',
    });
  });

  return (
    <Dashboard>
      <Container maxWidth="xl" sx={{ mt: 4, mb: 4 }}>
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">MAC Address</TableCell>
                <TableCell align="right">IP</TableCell>
                <TableCell align="right">Failures</TableCell>
                <TableCell align="right">Active</TableCell>
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
                  <TableCell align="right">{device.Failures}</TableCell>
                  <TableCell align="right">{device.Active ? <CheckIcon /> : <ClearIcon />}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </Dashboard>
  );
}
