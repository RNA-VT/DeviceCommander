import * as React from 'react'
import { useState } from 'react'
import Device from '../../utils/Device'
import { makeStyles, TextField, Button } from '@material-ui/core';
import SolenoidTable from './SolenoidTable'

const useStyles = makeStyles({
  title: {
    margin: 0,
  },
  buttonContainer: {
    marginTop: "10px"
  }
});

type DeviceProps = {
  device: Device,
  reload: () => Promise<void>
}

const DeviceForm = ({ device: device, reload }: DeviceProps) => {
  const classes = useStyles({})
  const [descriptionValue, setDescription] = useState(device.description)
  const [solenoids, setSolenoids] = useState(device.solenoids)

  const handleDeviceSave = (event: any) => {
    event.preventDefault()
    device.edit({
      description: descriptionValue
    }).then((data) => {
      reload()
    })
  }

  const handleSolenoidChange = (id: string, newValue: any, target: string) => {

    console.log('handleSolenoidChange', id, newValue, target)
    const index = solenoids.findIndex((solenoid) => {
      return solenoid.uid == id
    })

    console.log('targetIndex', index, solenoids[index])

    switch (target) {
      case 'name':
        solenoids[index].name = newValue
        break;
    }
  }

  const handleReset = () => {
    setDescription(device.description)
  }

  return (
    <>
      <form onSubmit={handleDeviceSave}>
        <p className={classes.title}><strong>ID:</strong> {device.id}</p>
        <p className={classes.title}><strong>URL:</strong> {device.host}:{device.port}</p>

        <TextField type="text" fullWidth
          className={classes.title}
          value={descriptionValue}
          onChange={(e) => setDescription(e.target.value)} />

        <div className={classes.buttonContainer}>
          <Button type="submit" variant="outlined">Submit</Button>
          <Button onClick={handleReset} variant="outlined">Clear Values</Button>
        </div>
      </form >

      <SolenoidTable solenoids={solenoids} isEdit={true} handleEdit={handleSolenoidChange} />
    </>

  )

}

export default DeviceForm