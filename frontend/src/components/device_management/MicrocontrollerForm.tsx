import * as React from 'react'
import { useState } from 'react'
import Microcontroller from '../../utils/Microcontroller'
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

type MicrocontrollerProps = {
  microcontroller: Microcontroller,
  reload: () => Promise<void>
}

const MicrocontrollerForm = ({ microcontroller, reload }: MicrocontrollerProps) => {
  const classes = useStyles({})
  const [descriptionValue, setDescription] = useState(microcontroller.description)
  const [solenoids, setSolenoids] = useState(microcontroller.solenoids)

  const handleMcSave = (event: any) => {
    event.preventDefault()
    microcontroller.edit({
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

    // setSolenoids(solenoids)

    console.log(solenoids[index])
    // 

  }

  const handleReset = () => {
    setDescription(microcontroller.description)
  }

  return (
    <>
      <form onSubmit={handleMcSave}>
        <p className={classes.title}><strong>ID:</strong> {microcontroller.id}</p>
        <p className={classes.title}><strong>URL:</strong> {microcontroller.host}:{microcontroller.port}</p>

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

export default MicrocontrollerForm