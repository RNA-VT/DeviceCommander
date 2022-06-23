import * as React from 'react'
import { useState } from 'react'
import { makeStyles, Button, Select, MenuItem, TableCell, TableRow, TextField, Grid } from '@material-ui/core'
import SettingsIcon from '@material-ui/icons/Settings'
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline';
import HighlightOffIcon from '@material-ui/icons/HighlightOff';
import Solenoid from '../../utils/Solenoid'
import { diff } from 'deep-object-diff';


type SolenoidRowProps = {
  solenoid: Solenoid,
  cellClasses: any,
  setIsEdit: React.Dispatch<React.SetStateAction<boolean>>
}

const useStyles = makeStyles({
  buttons: {
  }
})

const SolenoidEditRow = ({ solenoid, cellClasses, setIsEdit }: SolenoidRowProps) => {
  const [thisSolenoid, setSolenoid] = useState(solenoid)

  const setName = (text: string) => {
    const tmpSolenoid = new Solenoid({ ...thisSolenoid.getConfig(), 'Name': text }, {})
    setSolenoid(tmpSolenoid)
  }

  const setStatus = (event: React.ChangeEvent<{ value: unknown }>) => {
    const choice = event.target.value as string === 'true' ? true : false
    const tmpSolenoid = new Solenoid({ ...thisSolenoid.getConfig(), 'Enabled': choice }, {})
    setSolenoid(tmpSolenoid)
  }

  const handleCloseAndSave = () => {
    if (thisSolenoid != solenoid) {
      const diffData = diff(solenoid.getConfig(), thisSolenoid.getConfig())
      solenoid.edit(diffData).then((data) => {
        console.log(data)
        setIsEdit(false)
      })
    } else {
      setIsEdit(false)
    }

  }

  return (
    <>
      <TableRow key={solenoid.uid}>
        <TableCell className={cellClasses.cells}>{solenoid.uid}</TableCell>
        <TableCell className={cellClasses.cells}>
          <TextField type="text" fullWidth
            value={thisSolenoid.name}
            onChange={(e) => setName(e.target.value)} />
        </TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.headerPin}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.type}</TableCell>
        <TableCell className={cellClasses.cells}>
          <Select value={thisSolenoid.enabled}
            onChange={setStatus}>
            <MenuItem value={'true'}>Enabled</MenuItem>
            <MenuItem value={'false'}>Disabled</MenuItem>
          </Select>
        </TableCell>
        <TableCell className={cellClasses.cells}>
          <Grid container
            direction="column"
            justify="center"
            alignItems="center">
            <Button onClick={() => handleCloseAndSave()}>
              <CheckCircleOutlineIcon />
            </Button>
            <Button onClick={() => setIsEdit(false)}>
              <HighlightOffIcon />
            </Button>
          </Grid>

        </TableCell>
      </TableRow>
    </>
  )
}

export default SolenoidEditRow