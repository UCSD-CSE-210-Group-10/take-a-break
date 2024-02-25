import React from 'react';
import './EventsPage.css';
import logo from './UCSD-logo.png';
// import eventImage from './event-image.jpg';
import NavigationBar from './NavigationBar';

const EventsPage = () => {
  // Dummy data for event cards
  const events = [
    { id: 1, name: 'Event 1', date: '2024-02-20', time: '10:00 AM', organization: 'Organization A' },
    { id: 2, name: 'Event 2', date: '2024-02-21', time: '2:00 PM', organization: 'Organization B' },
    { id: 3, name: 'Event 3', date: '2024-02-22', time: '6:00 PM', organization: 'Organization C' },
  ];

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
            <a key={event.id} href={`/events/${event.id}`} className="event-card-link">
              <div className="event-card">
                <img src={logo} alt="Event" className="event-image" />
                <h3>{event.name}</h3>
                <p>
                  <span>{event.date}</span> | <span>{event.time}</span> 
                </p>
                <p>{event.organization}</p>
              </div>
            </a>
          ))}
        </div>
      </div>
    </div>
  );
}

export default EventsPage;
