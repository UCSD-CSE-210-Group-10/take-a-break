import React, { useState, useEffect } from "react";
import "./EventDetails.css";
import backButton from "./return-button.png";
import { Link, useParams } from "react-router-dom";
import dummyPoster from "./dummy-poster.png";
import NavigationBar from "./NavigationBar";

const EventDetails = () => {
	// State to handle RSVP button
	const [rsvpButtonText, setRsvpButtonText] = useState("RSVP");
	const [rsvpButtonDisabled, setRsvpButtonDisabled] = useState(false);

	const [event, setEvent] = useState([]);
	const [attendingFriends, setAttendingFriends] = useState([]);

	let { id } = useParams();
	console.log(id);

	useEffect(() => {
		
		// Function to fetch events from the API
		const fetchEventByID = async () => {
			try {
				const response = await fetch(`http://localhost:8080/events/${id}`);
				const data = await response.json();
				setEvent(data); // Assuming the API response contains an array of events
			} catch (error) {
				console.error("Error fetching events:", error);
			}
		};
		const fetchUserEvent = async () => {
			try {
				const jwtToken = localStorage.getItem('token');
				const response = await fetch(`http://localhost:8080/user_event/${jwtToken}/${id}`, {method: "GET"});
				const data = await response.json();
				console.log(data);
				if (data.email_id) {
					setRsvpButtonText("Going");
					setRsvpButtonDisabled(true);
				}
			} catch (error) {
				console.error("Error fetching user event:", error);
			}
		};
		const fetchAttendingFriends = async () => {
			try {
			const jwtToken = localStorage.getItem('token');
			  const response = await fetch(`http://localhost:8080/friends/attendance/${jwtToken}/${id}`);
			  if (!response.ok) {
				throw new Error('Failed to fetch data');
			  }
			  const data = await response.json();
			  console.log("Attending Friends Data:", data);
			  setAttendingFriends(data);
			} catch (error) {
			  console.error("Error fetching attending friends:", error);
			}
		  };
		  
	
        // Call the fetch functions
		fetchEventByID()
  		fetchUserEvent()
		fetchAttendingFriends()
	}, [id]); // Empty dependency array ensures the effect runs once when the component mounts


	const handleRsvpButtonClick = async () => {	
		try {
			const jwtToken = localStorage.getItem('token');
			const response = await fetch(`http://localhost:8080/user_event/${jwtToken}/${id}`, {method: "POST"});
	
			if (!response.ok) {
				throw new Error("Failed to RSVP");
			}	

			setRsvpButtonText("Going");
            setRsvpButtonDisabled(true);

		} catch (error) {
			console.error("Error RSVPing:", error);
		}
	};


	return (
		<div>
			<NavigationBar />
			<div className="event-details-container">
				<div className="back-button-container">
					<Link to="/events">
						<button className="back-button">
							<img src={backButton} className="back-png" alt="Back" />
						</button>
					</Link>
				</div>
				<div className="event-details-content">
					<div
						className="left-section-events"
						data-testid="left-section-events"
					>
						<div className="event-date-time">
							{new Date(event.date).toDateString()} |{" "}
							{new Date(event.time).toLocaleTimeString("en-US")}
						</div>
						<div className="event-info">
							<h1 className="event-name">{event.title}</h1>
							<button
								className={`rsvp-button ${rsvpButtonText === "Going" ? "going-button" : "rsvp-button"}`}
								onClick={handleRsvpButtonClick}
								disabled={rsvpButtonDisabled}
							>
								{rsvpButtonText}
							</button>
						</div>
						<div className="poster">
							<img src={dummyPoster} alt="dummy-poster"></img>
						</div>
						<div className="description">{event.description}</div>
					</div>

					<div
						className="right-section-events"
						data-testid="right-section-events"
					>
						<div className="friends-section">
							<p className="friends-p">
							<span className="friends-label">Friends Attending</span> 
							<br />
							{attendingFriends && attendingFriends.length > 0 ? (
								<span className="friends-s">
									{attendingFriends.length === 1
										? <><strong>{attendingFriends[0].name}</strong> is going.</>
										: attendingFriends.length === 2
										? <><strong>{attendingFriends[0].name}</strong> and <strong> {attendingFriends[1].name}</strong>  are going.</>
										: <><strong>{attendingFriends[0].name}</strong>, <strong>{attendingFriends[1].name}</strong>, and <strong>{attendingFriends.length - 2} others</strong> are going.</>}
								</span>
							) : (
								<span className="friends-s">No friends are attending this event.</span>
							)}
							</p>
							<br/>
						</div>
						<div className="details-section">
							<p className="event-details-p">
								<span className="label">Location</span>
								<br />
								<span className="info">{event.venue}</span>
							</p>
							<p className="event-details-p">
								<span className="label">Date and Time</span>
								<br />
								<span className="info">
									{new Date(event.date).toDateString()} |{" "}
									{new Date(event.time).toLocaleTimeString("en-US")}
								</span>
							</p>
							<p className="event-details-p">
								<span className="label">Event Fee</span>
								<br />
								<span className="info">Free</span>
							</p>
							<p className="event-details-p">
								<span className="label">Contact</span>
								<br />
								<span className="info">{event.contact}</span>
							</p>
							<p className="event-details-p">
								<span className="label">Audience</span>
								<br />
								<span className="info">Graduate and Professional students</span>
							</p>
							<p className="event-details-p">
								<span className="label">Event Host/Organization</span>
								<br />
								<span className="info">{event.host}</span>
							</p>
							<p className="event-details-p">
								<span className="label">Event Category</span>
								<br />
								<span className="info">{event.tags}</span>
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
};

export default EventDetails;
