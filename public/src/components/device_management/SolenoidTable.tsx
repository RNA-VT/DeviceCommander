import * as React from 'react'
import {
  Card,
  makeStyles,
  Grid,
  Button
} from '@material-ui/core'

import EditIcon from '@material-ui/icons/Edit'

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

const SolenoidTable = ({ solenoids }) => {
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
          solenoids.map((solenoid) => {
            return (
              <tr key={solenoid.UID}>
                <td className={classes.cells}>{solenoid.UID}</td>
                <td className={classes.cells}>{solenoid.Name}</td>
                <td className={classes.cells}>{solenoid.HeaderPin}</td>
                <td className={classes.cells}>{solenoid.Type}</td>
                <td className={classes.cells}>{solenoid.Enabled ? 'Enabled' : 'Disabled'}</td>
                <td className={classes.cells}><Button><EditIcon /></Button></td>
              </tr>
            )
          })
        }
      </table>
    </>
  )
}

export default SolenoidTable