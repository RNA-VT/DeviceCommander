import * as React from 'react'
import { render } from 'react-dom'
import { BrowserRouter as Router } from 'react-router-dom'
import { Provider } from 'unstated-typescript'
import { normalize } from 'polished'
import Routes from './Routes'

render(
    <Provider>
        <Router>
            <Routes />
        </Router>
    </Provider>,
    document.getElementById('app')
)
