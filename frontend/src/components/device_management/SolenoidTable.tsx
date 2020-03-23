import * as React from 'react'
import { makeStyles, Table, TableBody, TableCell, TableHead, TableRow } from '@material-ui/core'

import SolenoidRow from "./SolenoidRow"
import SolenoidEditRow from "./SolenoidEditRow"
import Solenoid from '../../utils/Solenoid'

const useStyles = makeStyles({
  table: {
    marginTop: '20px'
  },
  cells: {
    // padding: '10px',
  },
  cellHeaders: {
    // padding: '10px 10px 0px',
    paddingBottom: '10px',
    fontWeight: 'bold',
    fontSize: '1.2em'
  }
});

type SolenoidTableProps = {
  solenoids: Array<Solenoid>,
  isEdit: boolean,
  handleEdit: (id: string, newValue: any, target: string) => void
}

const SolenoidTable = ({ solenoids, isEdit, handleEdit }: SolenoidTableProps) => {
  const classes = useStyles({})
  let rows = solenoids.map((solenoid: Solenoid) => {
    return (
      <SolenoidRow
        key={solenoid.uid}
        solenoid={solenoid}
        cellClasses={classes.cells} />
    )
  })

  if (isEdit) {
    rows = solenoids.map((solenoid: Solenoid) => {
      return (
        <SolenoidEditRow
          key={solenoid.uid}
          solenoid={solenoid}
          cellClasses={classes.cells}
          handleEdit={handleEdit} />
      )
    })
  }


  return (
    <>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell className={classes.cellHeaders}>UID</TableCell>
            <TableCell className={classes.cellHeaders}>Name</TableCell>
            <TableCell className={classes.cellHeaders}>Pin</TableCell>
            <TableCell className={classes.cellHeaders}>Type</TableCell>
            <TableCell className={classes.cellHeaders}>Enabled</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows}
        </TableBody>


      </Table>
    </>
  )
}

export default SolenoidTable