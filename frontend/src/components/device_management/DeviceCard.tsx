import * as React from 'react'
import SolenoidTable from "./SolenoidTable"
import {
  Card,
  makeStyles,
  Grid
} from '@material-ui/core'

const useStyles = makeStyles({
  card: {
    minWidth: 275,
    padding: '20px',
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    margin: 0,
  },
  pos: {
    marginBottom: 12,
  },
});

const DeviceCard = (
  { children, microcontroller }: { children?: React.ReactNode, microcontroller: any }
) => {
  const classes = useStyles({})

  return (
    <Card className={classes.card}>
      <p className={classes.title}><strong>ID:</strong> {microcontroller.ID}</p>
      <p className={classes.title}><strong>URL:</strong> {microcontroller.Host}:{microcontroller.Port}</p>
      <p className={classes.title}><strong>Description:</strong> {microcontroller.Description}</p>
      <Grid container spacing={3}>
        <Grid item sm>
          <SolenoidTable solenoids={microcontroller.Solenoids} />
        </Grid>
      </Grid>
      {children}
    </Card >
  )
}

export default DeviceCard
