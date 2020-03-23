import * as React from 'react'
import { makeStyles, Table, TableBody, TableCell, TableHead, TableRow } from '@material-ui/core'

import SolenoidRow from "./SolenoidRow"
import Solenoid from '../../utils/Solenoid';

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
  isEdit: boolean
}

const SolenoidTable = ({ solenoids, isEdit }: SolenoidTableProps) => {
  const classes = useStyles({})

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
          {solenoids.map((solenoid: Solenoid) => {
            return (
              <SolenoidRow
                key={solenoid.uid}
                solenoid={solenoid}
                cellClasses={classes.cells} />
            )
          })}
        </TableBody>


      </Table>
    </>
  )
}

export default SolenoidTable