import "./App.css";

import EventsPage from "./EventsPage";
import CreateEvent from "./CreateEvent";
import Login from "./Login";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import EventDetails from "./EventDetails";

function App() {
  return (
    <div className="App" data-testid="app-head">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Login />}></Route>
          <Route path="/events" element={<EventsPage />}></Route>
          <Route path="/events/1" element={<EventDetails />}></Route>
          <Route path="/admin/events/create" element={<CreateEvent />}></Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
