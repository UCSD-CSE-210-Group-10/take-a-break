import React from 'react';
import './EventDetails.css';
import backButton from './return-button.png';
import dummyPoster from './dummy-poster.png';
import NavigationBar from './NavigationBar';

const EventDetails = () => {
    const dummyDescription = "Join us for a night of fun and games! Variety of card and board games to play! Bring your friends and/or have the opportunity to meet some new people!";

    
    return (
      <div> 
        <NavigationBar />
        <div className="event-details-container">
            <div className="back-button-container">
            <a href="/events">
                <button className="back-button"><img src={backButton} className="back-png" alt="Back" /></button>
                </a>
            </div>
            <div className="event-details-content">
                <div className="left-section" data-testid="left-section">
                    <div className="event-date-time">
                        Feb 23, 2024 | 6:00 - 9:00 PM
                    </div>
                    <div className="event-info">
                      <h1 className="event-name">GPSA Game Night</h1>
                      <button className="rsvp-button">RSVP</button>    
                    </div>
                    <div className="poster">
                        <img src={dummyPoster} alt="dummy-poster"></img>
                    </div>
                    <div className="description">
                        {dummyDescription}
                    </div>
                </div>
                <div className="right-section" data-testid="right-section">
                <div className="details-section">
                        <p className="event-details-p">
                            <span className="label">Location</span><br />
                            <span className="info">Dirty Bird @ Price Center Plaza</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Date and Time</span><br />
                            <span className="info">Friday | Feb. 23, 2024 | 6:00-9:00 PM</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Event Fee</span><br />
                            <span className="info">Free</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Contact</span><br />
                            <span className="info">gpsa.ucsd.edu</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Audience</span><br />
                            <span className="info">Graduate and Professional students</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Event Host/Organization</span><br />
                            <span className="info">GPSA</span>
                        </p>
                        <p className="event-details-p">
                            <span className="label">Event Category</span><br />
                            <span className="info">Graduate, Free Food, In-Person</span>
                        </p>
                    </div>
                </div>
            </div>
          </div>
        </div>
    );
}

export default EventDetails;
