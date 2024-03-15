import React, { useState, useEffect } from "react";
import { Carousel } from 'react-responsive-carousel';
import "react-responsive-carousel/lib/styles/carousel.min.css";
import './WellfarePage.css';
import { Link } from "react-router-dom";
import building from '../../images/SHS-building.png';
import staff from '../../images/Staff.png';
import contact from '../../images/Contact.png';
import hotline from '../../images/Hotline.png';
import NavigationBar from '../../components/nav_bar/NavigationBar';
import EventCard from '../EventCard'; // Import the EventCard component

const WellfarePage = ({ handleLogout }) => {
  const [events, setEvents] = useState([]);
  const [showMore, setShowMore] = useState(false); // State to manage showing more events

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const { hostname, protocol } = window.location;
        const jwtToken = localStorage.getItem('token');
        const response = await fetch(`${protocol}//${hostname}:8080/events/all/${jwtToken}`);
        const data = await response.json();
        if(data.error && data.error === "Auth Error") {
					handleLogout()
				}
        setEvents(data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    fetchEvents();
  }, [handleLogout]);

  // Function to handle showing more events
  const handleShowMore = () => {
    setShowMore(true);
  };

  // Filter events to show maximum of 6 events if showMore state is false
  const filteredEvents = showMore ? events : events.slice(0, 6).filter(event => event.tags.includes("Tag3"));

  return (
    <div>
      <NavigationBar handleLogout={handleLogout}/>
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
        <div className="section">
          <div className="left-column">
            <h2 className="section-title">Student Health Services</h2>
            <p className="contact-info">
              <span><strong> Contact Info:</strong></span><br/>
              <span> Appointments / Urgent Care - (858) 534-3300 </span><br/>
              <a href="https://mystudentchart.ucsd.edu/shs" target="_blank">Online - MyStudentChart</a>  
            </p>
            <p className="hours-locations">              
              <span> <strong>Hours & Locations:</strong></span><br/>
              <span> Days: Monday - Friday, Hours: 8:00 am - 4:00 pm</span><br/>
              <span> Library Walk, west of the Price Center, south of Geisel Library </span><br/>
              <span> * Offices closed during University holidays</span>
            </p>
          </div>
          <div className="right-column"> 
            <Carousel className="service-carousel" autoPlay={true} infiniteLoop={true} showThumbs={false} showStatus={false}>
              <div onClick={() => window.open('https://studenthealth.ucsd.edu/about/index.html', '_blank')}>
                <img src={building} alt="SHS Building" />
              </div>
              <div>
                <img src={staff} alt="Staff" />
              </div>
            </Carousel>
          </div>
        </div>   

        <div className="section">
          <div className="left-column">  
            <h2 className="section-title">Counselling and Psychological Services</h2>
            <p className="contact-info">
              <span><strong> Contact Info:</strong></span><br/>
              <span> Appointments / Urgent Care - (858) 534-3755</span><br/>
              <a href="https://mystudentchart.ucsd.edu/shs" target="_blank">Online - MyStudentChart</a>  
            </p>
            
            <p className="hours-locations">
              <span> <strong>Hours & Locations:</strong></span><br/>
              <span> Days: Monday - Friday, Hours: 8:00 am - 4:00 pm</span><br/>
              <span> Central Office & Urgent Care: Galbraith Hall 190</span><br/>
              <span> College & Satellite Offices: Locations Page</span><br/>
              <span> * Offices closed during University holidays</span>
            </p>
          </div>
          <div className="right-column"> 
            <Carousel className="service-carousel" autoPlay={true} infiniteLoop={true} showThumbs={false} showStatus={false}>
              <div onClick={() => window.open('https://caps.ucsd.edu/services/crisis.html', '_blank')}>
                  <img src={contact} alt="Contact" />
              </div>
              <div onClick={() => window.open('https://caps.ucsd.edu/services/crisis.html', '_blank')}>
                <img src={hotline} alt="Hotline" />
              </div>
            </Carousel>
          </div>
        </div>
        


      </div>
    </div>
  );
}

export default WellfarePage;
