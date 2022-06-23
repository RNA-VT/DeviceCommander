

import * as React from 'react'
import { HashRouter as Router, Routes, Route, Link } from 'react-router-dom'
import { render } from 'react-dom';
import Home from './pages/Home';
import './App.css';
import Settings from './pages/Settings';
import Devices from './pages/Devices';


export default class App extends React.Component {
  render() {
    return (
      <Router>
        <div>
          {/* <nav>
            <Link to="/">Home</Link>
            <Link to="/foo">Foo</Link>
            <Link to="/bar">Bar</Link>
          </nav> */}
          <Routes>
            <Route path="/" element={<Home/>} />
            <Route path="/devices" element={<Devices/>} />
            <Route path="/settings" element={<Settings/>} />
            
            {/* <Route path="/foo" element={Foo} />
            <Route path="/bar" element={Bar} /> */}
          </Routes>
        </div>
      </Router>
    );
  }
}

render(<App />, document.getElementById('root'));
