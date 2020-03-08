import * as React from 'react'
import { useState } from 'react'
import SolenoidTable from './SolenoidTable'
import {
  Card,
  makeStyles,
  Grid,
  Button
} from '@material-ui/core'
import SettingsIcon from '@material-ui/icons/Settings'

const useStyles = makeStyles({
  card: {
    minWidth: 275,
    padding: '20px',
    position: 'relative'
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    margin: 0,
  },
  settingsButton: {
    position: 'absolute',
    right: '20px',
    top: '15px'
  },
  pos: {
    marginBottom: 12,
  },
});

type DeviceCardProps = {
  children?: React.ReactNode,
  microcontroller: any
}

const DeviceCard = ({ children, microcontroller }: DeviceCardProps) => {
  const classes = useStyles({})
  const [isEdit, setIsEdit] = useState(false)

  return (
    <Card className={classes.card}>
      <Button className={classes.settingsButton}>
        <SettingsIcon />
      </Button>

      <p className={classes.title}><strong>ID:</strong> {microcontroller.ID}</p>
      <p className={classes.title}><strong>URL:</strong> {microcontroller.Host}:{microcontroller.Port}</p>
      <p className={classes.title}><strong>Description:</strong> {microcontroller.Description}</p>
      <Grid container spacing={3}>
        <Grid item sm>
          <SolenoidTable solenoids={microcontroller.Solenoids} isEdit={isEdit} />
        </Grid>
      </Grid>
      {children}
    </Card >
  )
}

export default DeviceCard
