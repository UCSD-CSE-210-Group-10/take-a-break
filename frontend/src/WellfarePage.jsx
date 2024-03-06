import React, { useState, useEffect } from "react";
import './WellfarePage.css';
import logo from './UCSD-logo.png';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard'; // Import the EventCard component

const WellfarePage = () => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);   // State to store selected tags
  const [showMore, setShowMore] = useState(false); // State to manage showing more events

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

  // Function to handle showing more events
  const handleShowMore = () => {
    setShowMore(true);
  };

  // Filter events to show maximum of 6 events if showMore state is false
  const filteredEvents = showMore ? events : events.slice(0, 6);

  return (
    <div>
      <NavigationBar />
      <div className="events_container">
        <h1 className="title">Health and Wellfare @ UC San Diego</h1>
        <div className="event-cards">
          {filteredEvents.map(event => (
            <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
              <EventCard event={event} />
            </Link>
          ))}
        </div>
        {!showMore && events.length > 6 && (
          <button onClick={handleShowMore} className="show-more-button">Show More</button>
        )}
      </div>
    </div>
  );
}

export default WellfarePage;
