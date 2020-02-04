import * as React from 'react'
import { Card, makeStyles } from '@material-ui/core'

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
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
});

const DeviceCard = (props) => {
  const classes = useStyles({})

  return (
    <Card className={classes.card}>
      {props.children}
    </Card>
  )
}

export default DeviceCard
