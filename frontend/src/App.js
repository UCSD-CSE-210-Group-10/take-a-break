import './App.css';
import Login from './Login';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
  return (
    <div className="App" data-testid="app-head">
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<Login/>}>
            </Route>
          </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;
