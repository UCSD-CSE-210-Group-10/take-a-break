import "./Main.css";

import EventsPage from "./events/EventsPage";
import CreateEvent from "./events/CreateEvent";
import Login from "./auth/Login";
import { Routes, Route } from "react-router-dom";
import EventDetails from "./events/EventDetails";
import Friends from "./friends/Friends";
import WellfarePage from "./events/wellfare_events/WellfarePage";
import UserProfile from "./user/UserProfile";

function Main() {
  const handleLogout = () => {
    // Delete the token from local storage
    localStorage.removeItem('token');
    
    const { hostname, protocol } = window.location;
    window.location.href = `${protocol}//${hostname}:3000`;
  };
  
  return (
    <div>
        <Routes>
          <Route path="/" element={<Login />}></Route>
          <Route path="/events" element={<EventsPage handleLogout = {handleLogout}/>}></Route>
          <Route path="/events/:id" element={<EventDetails handleLogout = {handleLogout}/>}></Route>
          <Route path="/friends" element={<Friends handleLogout = {handleLogout}/>}></Route>
          <Route path="/health" element={<WellfarePage handleLogout = {handleLogout}/>}></Route>
          <Route path="/admin/events/create" element={<CreateEvent />}></Route>
          <Route path="/profile" element={<UserProfile handleLogout = {handleLogout}/>}></Route>
        </Routes>
    </div>
  );
}

export default Main;
