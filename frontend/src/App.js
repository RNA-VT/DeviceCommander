import * as React from 'react'
import { render } from 'react-dom'
import { BrowserRouter as Router } from 'react-router-dom'
import { Provider } from 'unstated-typescript'
import { normalize } from 'polished'
import Routes from './Routes'

function App() {
  return (
    <Provider>
      <Router>
        <Routes />
      </Router>
    </Provider>
  );
}

export default App;
