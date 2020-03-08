import * as React from 'react'
import { useState } from 'react'
import SolenoidTable from './SolenoidTable'
import MicrocontrollerForm from './MicrocontrollerForm'
import {
  Card,
  makeStyles,
  Grid,
  Button,
} from '@material-ui/core'
import SettingsIcon from '@material-ui/icons/Settings'
import Microcontroller from '../../utils/Microcontroller'

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
  microcontroller: Microcontroller,
  reload: () => Promise<void>
}

const DeviceCard = ({ children, microcontroller, reload }: DeviceCardProps) => {
  const classes = useStyles({})
  const [isEdit, setIsEdit] = useState(false)
  let basicInfo = null

  if (isEdit) {
    basicInfo = (
      <MicrocontrollerForm microcontroller={microcontroller} reload={reload} />
    )
  } else {
    basicInfo = (
      <>
        <p className={classes.title}><strong>ID:</strong> {microcontroller.id}</p>
        <p className={classes.title}><strong>URL:</strong> {microcontroller.host}:{microcontroller.port}</p>
        <p className={classes.title}><strong>Description:</strong> {microcontroller.description}</p>
      </>
    )
  }

  return (
    <Card className={classes.card}>
      <Button className={classes.settingsButton} onClick={() => setIsEdit(!isEdit)}>
        <SettingsIcon />
      </Button>
      {basicInfo}


      <Grid container spacing={3}>
        <Grid item sm>
          <SolenoidTable solenoids={microcontroller.solenoids} isEdit={isEdit} />
        </Grid>
      </Grid>
      {children}
    </Card >
  )
}

export default DeviceCard
