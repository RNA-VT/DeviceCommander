import * as React from 'react'
import { TableCell, TableRow, TextField } from '@material-ui/core'
import Solenoid from '../../utils/Solenoid'

type SolenoidRowProps = {
  solenoid: Solenoid,
  cellClasses: any,
  handleEdit: (id: string, newValue: any, target: string) => void
}

const SolenoidEditRow = ({ solenoid, cellClasses, handleEdit }: SolenoidRowProps) => {
  return (
    <>
      <TableRow key={solenoid.uid}>
        <TableCell className={cellClasses.cells}>{solenoid.uid}</TableCell>
        <TableCell className={cellClasses.cells}>
          <TextField type="text" fullWidth
            value={solenoid.name}
            onChange={(e) => handleEdit(solenoid.uid, e.target.value, 'name')} />
        </TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.headerPin}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.type}</TableCell>
        <TableCell className={cellClasses.cells}>{solenoid.enabled ? 'Enabled' : 'Disabled'}</TableCell>
      </TableRow>
    </>
  )
}

export default SolenoidEditRow