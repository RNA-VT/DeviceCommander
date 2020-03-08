import * as React from 'react'
import { makeStyles } from '@material-ui/core'

import SolenoidRow from "./SolenoidRow"

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
  solenoids: Array<any>,
  isEdit: boolean
}

const SolenoidTable = ({ solenoids, isEdit }: SolenoidTableProps) => {
  const classes = useStyles({})

  return (
    <>
      <table className={classes.table}>
        <tr>
          <th className={classes.cellHeaders}>UID</th>
          <th className={classes.cellHeaders}>Name</th>
          <th className={classes.cellHeaders}>Pin</th>
          <th className={classes.cellHeaders}>Type</th>
          <th className={classes.cellHeaders}>Enabled</th>
        </tr>
        {
          solenoids.map((solenoid: any) => {
            return (
              <SolenoidRow
                solenoid={solenoid}
                cellClasses={classes.cells} />
            )
          })
        }
      </table>
    </>
  )
}

export default SolenoidTable