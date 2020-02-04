import * as React from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  wrapper: {

  }
});

const Wrapper = props => {
  const classes = useStyles({})
  return (
    <div className={classes.wrapper}>
      {props.children}
    </div>
  )
}

export default Wrapper
