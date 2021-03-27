import * as React from 'react'
import { useState } from 'react'
import SolenoidTable from './SolenoidTable'
import DeviceForm from './DeviceForm'
import {
  Card,
  makeStyles,
  Grid,
  Button,
} from '@material-ui/core'
import SettingsIcon from '@material-ui/icons/Settings'
import Device from '../../utils/Device'

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
  device: Device,
  reload: () => Promise<void>
}

const DeviceCard = ({ children, device: device, reload }: DeviceCardProps) => {
  const classes = useStyles({})
  const [isEdit, setIsEdit] = useState(false)
  let basicInfo = null

  if (isEdit) {
    basicInfo = (
      <DeviceForm device={device} reload={reload} />
    )
  } else {
    basicInfo = (
      <>
        <Grid
          container
          direction="column">
          <p className={classes.title}><strong>ID:</strong> {device.id}</p>
          <p className={classes.title}><strong>URL:</strong> {device.host}:{device.port}</p>
          <p className={classes.title}><strong>Description:</strong> {device.description}</p>
        </Grid>

        <Grid container spacing={3}>
          <Grid item sm>
            <SolenoidTable solenoids={device.solenoids} isEdit={isEdit} handleEdit={(id: string, newValue: any, target: string) => { }} />
          </Grid>
        </Grid>
      </>
    )
  }

  return (
    <Card className={classes.card}>
      <Button className={classes.settingsButton} onClick={() => setIsEdit(!isEdit)}>
        <SettingsIcon />
      </Button>

      {basicInfo}

      {children}
    </Card >
  )
}

export default DeviceCard
