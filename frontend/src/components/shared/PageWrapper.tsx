import * as React from 'react'
import { makeStyles } from '@material-ui/core'
import Layout from './Layout'

const useStyles = makeStyles({
  wrapper: {

  }
});

const Wrapper = ({ children }: { children: React.ReactNode }) => {
  const classes = useStyles({})
  return (
    <Layout>
      <div className={classes.wrapper}>
        {children}
      </div>
    </Layout>

  )
}

export default Wrapper
