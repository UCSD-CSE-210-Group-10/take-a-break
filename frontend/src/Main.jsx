import "./Main.css";

import EventsPage from "./EventsPage";
import CreateEvent from "./CreateEvent";
import Login from "./Login";
import { Routes, Route } from "react-router-dom";
import EventDetails from "./EventDetails";
import Friends from "./Friends";
import WellfarePage from "./WellfarePage";
import UserProfile from "./UserProfile";

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
