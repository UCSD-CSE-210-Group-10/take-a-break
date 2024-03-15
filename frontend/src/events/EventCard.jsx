// EventCard.js

import React from 'react';

import './EventCard.css'; // Create a CSS file for styling
import logo from '../images/UCSD-logo.png';


const EventCard = ({ event }) => {
  return (
    <div className="event-card" data-testid="event-card">
      <div className="event-image">
        {/* <img src={event.image} alt={event.title} /> */}
        <img src={logo} alt="Event" className="event-image" />
      </div>
      <div className="event-header">
        <h3>{event.title}</h3>
        <p>{new Date(event.date).toDateString()} | {new Date(event.time).toLocaleTimeString("en-US")}</p>
        <p>{event.host}</p>
      </div>
    </div>
  );
};

export default EventCard;