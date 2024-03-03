import React, { useState, useEffect } from "react";
import './EventDetails.css';
import backButton from './return-button.png';
import { Link, useParams } from 'react-router-dom';
import dummyPoster from './dummy-poster.png';
import NavigationBar from './NavigationBar';

const EventDetails = () => {

    const [event, setEvent] = useState([]);

    let {id} = useParams();
    console.log(id);

    useEffect(() => {
        // Function to fetch events from the API
        const fetchEventByID = async () => {
        try {
            const response = await fetch(`http://localhost:8080/events/${id}`);
            const data = await response.json();
            setEvent(data); // Assuming the API response contains an array of events
        } catch (error) {
            console.error('Error fetching events:', error);
        }
        };

        // Call the fetchEvents function
        fetchEventByID();
    }, [id]); // Empty dependency array ensures the effect runs once when the component mounts

    return (
      <div> 
        <NavigationBar />
        <div className="event-details-container">
            <div className="back-button-container">
                <Link to="/events">
                    <button className="back-button"><img src={backButton} className="back-png" alt="Back" /></button>
                </Link>
            </div>
            <div className="event-details-content">
                <div className="left-section" data-testid="left-section">
                    <div className="event-date-time">
                        {new Date(event.date).toDateString()} | {new Date(event.time).toLocaleTimeString("en-US")}
                    </div>
                    <div className="event-info">
                      <h1 className="event-name">{event.title}</h1>
                      <button className="rsvp-button">RSVP</button>    
                    </div>
                    <div className="poster">
                        <img src={dummyPoster} alt="dummy-poster"></img>
                    </div>
                    <div className="description">
                        {event.description}
                    </div>
                </div>
                <div className="right-section" data-testid="right-section">
                <div className="details-section">
                        <p className="event-details-p">
                            <span className="label">Location</span><br />
                            <span className="info">{event.venue}</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Date and Time</span><br />
                            <span className="info">{new Date(event.date).toDateString()} | {new Date(event.time).toLocaleTimeString("en-US")}</span>
                        </p>
                        {/* <p className="event-details-p">
                            <span className="label">Event Fee</span><br />
                            <span className="info">Free</span>
                        </p> */}
                        <p className="event-details-p">
                            <span className="label">Contact</span><br />
                            <span className="info">{event.contact}</span>
                        </p>
                        {/* <p className="event-details-p">
                            <span className="label">Audience</span><br />
                            <span className="info">Graduate and Professional students</span>
                        </p> */}
                        <p className="event-details-p">
                            <span className="label">Event Host/Organization</span><br />
                            <span className="info">{event.host}</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Event Category</span><br />
                            <span className="info">{event.tags}</span>
                        </p>
                    </div>
                </div>
            </div>
          </div>
        </div>
    );
}

export default EventDetails;
