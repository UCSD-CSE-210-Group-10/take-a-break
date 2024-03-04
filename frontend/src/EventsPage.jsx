import React, { useState, useEffect } from "react";
import './EventsPage.css';
import logo from './UCSD-logo.png';
import { Link } from "react-router-dom";
// import eventImage from './event-image.jpg';
import NavigationBar from './NavigationBar';

const EventsPage = () => {
  // Dummy data for event cards

  const [events, setEvents] = useState([]);

  useEffect(() => {
    // Function to fetch events from the API
    const fetchEvents = async () => {
      try {
        const response = await fetch('http://localhost:8080/events');
        const data = await response.json();
        setEvents(data); // Assuming the API response contains an array of events
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    // Call the fetchEvents function
    fetchEvents();
  }, []); // Empty dependency array ensures the effect runs once when the component mounts


  // Functionality for searching events
  const handleSearch = (event) => {
    // Implementation for searching events
    console.log(event.target.value);
  };

  return (
    <div>
      <NavigationBar />
      <div className="events-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar">
            <input type="text" placeholder="Search Event" onChange={handleSearch} className="search-input" />
            <select className="tags-dropdown" >
              <option value=""> Tags </option> {/* Empty default option */}
              <option value="Physical Wellness">Physical Wellness</option>
              <option value="Cultural Exchange">Cultural Exchange</option>
              <option value="LGBTQ+">LGBTQ+</option>
              <option value="Arts Entertainment">Arts/Entertainment</option>
              <option value="Graduate">Graduate</option>
              <option value="Undergraduate">Undergraduate</option>
              <option value="Virtual">Virtual</option>
              <option value="In Person">In Person</option>
              <option value="Free Food">Free Food</option>
            </select>
          </div>
        </div>
        <div className="event-cards">
          {events.map(event => (
            <div>
            <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
              <div className="event-card">
                <img src={logo} alt="Event" className="event-image" />
                <h3>{event.title}</h3>
                <p>
                  <span>{new Date(event.date).toDateString()}</span> | <span>{new Date(event.time).toLocaleTimeString("en-US")}</span> 
                </p>
                <p>{event.host}</p>
              </div>
              </Link>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default EventsPage;
