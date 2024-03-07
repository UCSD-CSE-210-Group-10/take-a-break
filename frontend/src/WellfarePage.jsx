import React, { useState, useEffect } from "react";
import './WellfarePage.css';
import logo from './UCSD-logo.png';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard'; // Import the EventCard component

const WellfarePage = () => {
  const [events, setEvents] = useState([]);
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
  const filteredEvents = showMore ? events : events.slice(0, 6).filter(event => event.tags.includes("Tag3"));

  return (
    <div>
      <NavigationBar />
      <div className="events_container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <h1 className="title">Health and Wellfare @ UC San Diego</h1>
          <h1 className="subtitle">UC San Diego is dedicated to supporting the well-being and academic achievements of every student. </h1>
          <h2 className="upcoming-events"> Upcoming Events </h2>

          <div className="event-cards">
            {filteredEvents.map(event => (
              <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
                <EventCard event={event} />
              </Link>
            ))}
          </div>
          {!showMore && events.filter(event => event.tags.includes("tag4")).length > 6 && (
            <button onClick={handleShowMore} className="show-more-button">Show More</button>
          )}
        </div>
      </div>
      <div className="wellfare-info">
        <div className="left-column">
          <div className="section">
            <h2 className="section-title">Student Health Services</h2>
            <p className="contact-info">
              <span><strong> Contact Info:</strong></span><br/>
              <span> Appointments / Urgent Care - (858) 534-3755</span><br/>
              <span> Online - MyStudentChart</span>
            </p>
            <p className="hours-locations">
              <span> <strong>Hours & Locations:</strong></span><br/>
              <span> Days: Monday - Friday, Hours: 8:00 am - 4:00 pm</span><br/>
              <span> Central Office & Urgent Care: Galbraith Hall 190</span><br/>
              <span> College & Satellite Offices: Locations Page</span><br/>
              <span> * Offices closed during University holidays</span>
            </p>
          </div>
          <div className="section">
            <h2 className="section-title">Counselling and Psychological Services</h2>
            <p className="contact-info">
              <span><strong> Contact Info:</strong></span><br/>
              <span> Appointments / Urgent Care - (858) 534-3755</span><br/>
              <span> Online - MyStudentChart</span>
            </p>
            <p className="hours-locations">
              <span> <strong>Hours & Locations:</strong></span><br/>
              <span> Days: Monday - Friday, Hours: 8:00 am - 4:00 pm</span><br/>
              <span> Central Office & Urgent Care: Galbraith Hall 190</span><br/>
              <span> College & Satellite Offices: Locations Page</span><br/>
              <span> * Offices closed during University holidays</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default WellfarePage;
