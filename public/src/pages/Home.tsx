import * as React from 'react'
import { Subscribe } from 'unstated-typescript'
import CounterContainer from '../containers/CounterContainer'
import Wrapper from '../components/shared/PageWrapper'
import PageHeader from '../components/shared/PageHeader'
import Button from '@material-ui/core/Button';

const Home = () => (
  <Wrapper>
    <PageHeader />
    <h1>Welcome to LightDash</h1>
    <Subscribe to={[CounterContainer]}>
      {counter => (
        <div>
          <h1 style={{ fontSize: '10rem', margin: '5% 0' }}>{counter.state.count}</h1>
          <Button
            variant="contained"
            onClick={() => counter.increment()}>Increment</Button>
          <Button
            variant="contained"
            onClick={() => counter.decrement()} text='Decrement'>Decrement</Button>
        </div>
      )}
    </Subscribe>
  </Wrapper>
)

export default Home
