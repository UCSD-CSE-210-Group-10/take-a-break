import React, { useState, useEffect } from 'react';
import Modal from 'react-modal';
import "./RequestModal.css";
import "react-notifications/lib/notifications.css";
import { NotificationManager } from "react-notifications";

const RequestModal = ({ isOpen, onRequestClose, jwtToken, handleLogout }) => {
  const [requests, setRequests] = useState([]);

  const acceptRequest = async (requestId) => {
    const jwtToken = localStorage.getItem('token');
    // Implement logic to accept the friend request
    try {
      const { hostname, protocol } = window.location;
      const response = await fetch(`${protocol}//${hostname}:8080/friends/request/accept/${jwtToken}`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email_id: `${requestId}`  }),
			});
	
			if (!response.ok) {
        NotificationManager.error("Can't Accept Request","", 10000);
				throw new Error("Failed to Accept Request");
			}

      NotificationManager.success("Accepted Request","", 10000);
      setRequests((prevRequests) => prevRequests.filter((request) => request.email_id !== requestId));
    } catch (error) {
      console.error('Error Accepting request:', error);
    }
  };

  const ignoreRequest = async (requestId) => {
    const jwtToken = localStorage.getItem('token');
    // Implement logic to accept the friend request
    try {
      const { hostname, protocol } = window.location;
      const response = await fetch(`${protocol}//${hostname}:8080/friends/request/ignore/${jwtToken}`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email_id: `${requestId}`  }),
			});
	
			if (!response.ok) {
        NotificationManager.error("Can't Ignore Request","", 10000);
				throw new Error("Failed to Ignore Request");
			}

      NotificationManager.success("Ignored Request","", 10000);
      setRequests((prevRequests) => prevRequests.filter((request) => request.email_id !== requestId));
    } catch (error) {
      console.error('Error Ignoring request:', error);
    }
  };

  useEffect(() => {
    const jwtToken = localStorage.getItem('token');
    // Fetch requests from the backend API
    const fetchRequests = async () => {
      try {
        const { hostname, protocol } = window.location;
        const response = await fetch(`${protocol}//${hostname}:8080/friends/request/get/${jwtToken}`);
        const data = await response.json();
        if(data.error && data.error === "Auth Error") {
					handleLogout()
				}
        setRequests(data); // Assuming the API response contains an array of requests
      } catch (error) {
        console.error('Error fetching requests:', error);
      }
    };

    // Call the fetchRequests function when the modal opens
    if (isOpen) {
      fetchRequests();
    }
  }, [isOpen, handleLogout]);

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onRequestClose}
      contentLabel="Friend Requests"
    >
      <div className='request-modal'>
        <h2>Friend Requests</h2>
        <button onClick={onRequestClose} className='closeButton'>Close</button>
      </div>
      <hr/>
      <ul>
        {requests && requests.map((request) => (
          <li className='requests-list' key={request.email_id}>
            <div>{request.name} has sent you a friend request.</div>
            <div>
            <button className="modalButton" onClick={() => acceptRequest(request.email_id)}>Accept</button>
            <button className="modalButton" onClick={() => ignoreRequest(request.email_id)}>Ignore</button>
            </div>
            
          </li>
        ))}
      </ul>
      
    </Modal>
  );
};

export default RequestModal;