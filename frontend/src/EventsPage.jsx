import React, { useState, useEffect } from "react";
import './EventsPage.css';
import logo from './UCSD-logo.png';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';

const EventsPage = () => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);   // State to store selected tags

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await fetch('http://localhost:8080/events');
        const data = await response.json();
        setEvents(data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    fetchEvents();
  }, []);

  const handleSearch = (event) => {
    
    // Implementation for searching events
    console.log(event.target.value);

    const selectedOptions = Array.from(event.target.selectedOptions).map(option => option.value);
    setSelectedTags(selectedOptions); // Update selected tags state
  };

  // Function to filter events based on selected tags
  const filteredEvents = selectedTags.length
    ? events.filter(event => selectedTags.every(tag => event.tags.includes(tag)))
    : events;

  return (
    <div>
      <NavigationBar />
      <div className="events-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar">
            <input type="text" placeholder="Search Event" onChange={handleSearch} className="search-input" />
            <select className="tags-dropdown" multiple onChange={handleSearch}>
              <option value=""> Tags </option>
              <option value="Tag1">Tag 1</option>
              <option value="Tag2">Tag 2</option>
              <option value="Tag3">Tag 3</option>
              <option value="Tag4">Tag 4</option>
              <option value="Tag5">Tag 5</option>
              {/* <option value="Physical Wellness">Physical Wellness</option>
              <option value="Cultural Exchange">Cultural Exchange</option>
              <option value="LGBTQ+">LGBTQ+</option>
              <option value="Arts Entertainment">Arts/Entertainment</option>
              <option value="Graduate">Graduate</option>
              <option value="Undergraduate">Undergraduate</option>
              <option value="Virtual">Virtual</option>
              <option value="In Person">In Person</option>
              <option value="Free Food">Free Food</option> */}
            </select>
          </div>
        </div>
        <div className="event-cards">
          {filteredEvents.map(event => (
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
          ))}
        </div>
      </div>
    </div>
  );
}

export default EventsPage;
