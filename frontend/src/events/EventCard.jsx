// EventCard.js

import React from 'react';
import './EventCard.css'; // Create a CSS file for styling


const EventCard = ({ event }) => {
  console.log(event);
  return (
    
    <div className="event-card" data-testid="event-card">
      <div className="event-image-container">
        {/* <img src={event.image} alt={event.title} /> */}
        <img src={event.imagepath} alt="Event" className="event-image" />
      </div>
      <div className="event-header">
        <h3>{event.title}</h3>
        
        <p>{new Date(event.date.substring(0, event.date.length-1).toLocaleString('en-US')).toDateString()} | {new Date(event.time.substring(0, event.time.length-1)).toLocaleTimeString("en-US")}</p>
        <p>{event.host}</p>
      </div>
    </div>
  );
};

export default EventCard;