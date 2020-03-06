import * as React from 'react'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles({
  wrapper: {

  }
});

const Wrapper = ({ children }: { children: React.ReactNode }) => {
  const classes = useStyles({})
  return (
    <div className={classes.wrapper}>
      {children}
    </div>
  )
}

export default Wrapper
