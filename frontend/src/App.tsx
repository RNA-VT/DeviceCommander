import * as React from 'react';
import { HashRouter as Router, Routes, Route } from 'react-router-dom';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Home from './pages/Home';
import './App.css';
import Settings from './pages/Settings';
import Devices from './pages/Devices';

const mdTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

export default function App() {
  return (
    <ThemeProvider theme={mdTheme}>
      <CssBaseline />
      <Router>
        <div>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/devices" element={<Devices />} />
            <Route path="/settings" element={<Settings />} />
          </Routes>
        </div>
      </Router>
    </ThemeProvider>
  );
}

// export default class App extends React.Component {
//   render() {
//     return (
//       <ThemeProvider theme={mdTheme}>
//         <CssBaseline />
//         <Router>
//           <div>
//             <Routes>
//               <Route path="/" element={<Home />} />
//               <Route path="/devices" element={<Devices />} />
//               <Route path="/settings" element={<Settings />} />
//             </Routes>
//           </div>
//         </Router>
//       </ThemeProvider>
//     );
//   }
// }
