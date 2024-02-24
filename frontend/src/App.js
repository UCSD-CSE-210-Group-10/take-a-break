import './App.css';

import EventsPage from './EventsPage';
// import CreateEvent from './CreateEvent';
import CreateEvent from './CreateEvent';
import Login from './Login';
import { BrowserRouter, Routes, Route } from 'react-router-dom';


function App() {
  return (
    <div className="App" data-testid="app-head">
      <EventsPage/>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<Login/>}></Route>
            <Route path='/admin/events/create' element={<CreateEvent/>}></Route>
          </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;