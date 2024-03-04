import "./Main.css";

import EventsPage from "./EventsPage";
import CreateEvent from "./CreateEvent";
import Login from "./Login";
import { Routes, Route } from "react-router-dom";
import EventDetails from "./EventDetails";
import EventDetailsSample from "./EventDetailsSample";

function Main() {
  return (
    <div>
        <Routes>
          <Route path="/" element={<Login />}></Route>
          <Route path="/events" element={<EventsPage />}></Route>
          <Route path="/events/:id" element={<EventDetails />}></Route>
          <Route path="/events/sampleEvent" element={<EventDetailsSample />}></Route>
          <Route path="/admin/events/create" element={<CreateEvent />}></Route>
        </Routes>
    </div>
  );
}

export default Main;
