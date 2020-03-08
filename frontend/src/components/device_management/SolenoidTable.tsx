import * as React from 'react'
import { makeStyles } from '@material-ui/core'

import SolenoidRow from "./SolenoidRow"
import Solenoid from '../../utils/Solenoid';

const useStyles = makeStyles({
  table: {
    marginTop: '20px'
  },
  cells: {
    padding: '10px',
  },
  cellHeaders: {
    padding: '10px 10px 0px',
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
      <table className={classes.table}>
        <thead>
          <tr>
            <th className={classes.cellHeaders}>UID</th>
            <th className={classes.cellHeaders}>Name</th>
            <th className={classes.cellHeaders}>Pin</th>
            <th className={classes.cellHeaders}>Type</th>
            <th className={classes.cellHeaders}>Enabled</th>
          </tr>
        </thead>
        <tbody>
          {solenoids.map((solenoid: Solenoid) => {
            return (
              <SolenoidRow
                key={solenoid.uid}
                solenoid={solenoid}
                cellClasses={classes.cells} />
            )
          })}
        </tbody>


      </table>
    </>
  )
}

export default SolenoidTable