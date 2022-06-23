import * as React from 'react'
import { useState } from 'react'
import SettingsIcon from '@material-ui/icons/Settings'
import { Button, TableCell, TableRow } from '@material-ui/core'
import Solenoid from '../../utils/Solenoid'
import SolenoidEditRow from './SolenoidEditRow'

type SolenoidRowProps = {
  solenoid: Solenoid,
  cellClasses: any
}

const SolenoidRow = ({ solenoid, cellClasses }: SolenoidRowProps) => {
  const [isEdit, setIsEdit] = useState(false)

  if (isEdit) {
    return (
      <SolenoidEditRow solenoid={solenoid}
        cellClasses={cellClasses}
        setIsEdit={setIsEdit} />
    )
  }

  return (
    <>
      <TableRow key={solenoid.uid}>
        <TableCell className={cellClasses.cells}>{solenoid.uid}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.name}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.headerPin}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.type}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.enabled ? 'Enabled' : 'Disabled'}</TableCell>
        <TableCell className={cellClasses.cells}>
          <Button onClick={() => setIsEdit(!isEdit)}>
            <SettingsIcon />
          </Button>
        </TableCell>
      </TableRow>
    </>
  )
}

export default SolenoidRow