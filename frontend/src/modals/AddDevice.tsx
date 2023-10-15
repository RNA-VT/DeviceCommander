import React from 'react';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import { RJSFSchema } from '@rjsf/utils';
import validator from '@rjsf/validator-ajv8';
import Form from '@rjsf/mui';
import AddDeviceMethod, { NewDeviceData } from '../api/method/AddDevice';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

export default function AddDeviceModal(props: { open: boolean, handleClose: () => void }) {
  const { open, handleClose } = props;

  const schema: RJSFSchema = {
    title: 'Add New Device',
    type: 'object',
    required: ['host', 'port'],
    properties: {
      name: { type: 'string', title: 'Name', default: 'Valiant Device #1' },
      host: { type: 'string', title: 'Host', default: '192.168.1.45' },
      port: { type: 'integer', title: 'Port', default: 8001 },
    },
  };

  const handleSubmit = async (data: any) => {
    console.log('Submitted', data);
    const { formData } : { formData: NewDeviceData } = data;
    console.log(formData);

    const newDeviceMethod = new AddDeviceMethod('http://localhost:8001', formData);

    const newDevice = await newDeviceMethod.do();
    console.log(newDevice);
    if (formData.host && formData.port) {
      console.log(`data is valid host and port ${formData.host}:${formData.port}`);
      handleClose();
    }
  };

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
    >
      <Box sx={style}>
        <Form
          schema={schema}
          validator={validator}
          onSubmit={handleSubmit}
        />
      </Box>
    </Modal>
  );
}
