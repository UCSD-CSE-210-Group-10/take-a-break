import React from 'react';
import './FriendCard.css'; // Create a CSS file for stylings

const FriendCard = ( {friend, updateSentRequest} ) => {

  const sendRequest = async (requestId) => {
    const jwtToken = localStorage.getItem('token');
    // Implement logic to accept the friend request
    try {
      const { hostname, protocol } = window.location;
      const response = await fetch(`${protocol}//${hostname}:8080/friends/request/send/${jwtToken}`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email_id: `${requestId}`  }),
			});
	
			if (!response.ok) {
				throw new Error("Failed to Send Request");
			}

      updateSentRequest(friend);
    } catch (error) {
      console.error('Error Sendings request:', error);
    }
  };

  return (
    <div className="friend-card" data-testid="friend-card">
      <div className="friend-image">
        <img src={friend.avatar} alt={friend.name} />
      </div>
      <div className="friend-info">
        <div className="friend-header">
          <h3>{friend.name}</h3>
          
          {(friend.hasOwnProperty('has_sent_request') && friend.has_sent_request === '0') && <button onClick={() => sendRequest(friend.email)} className="add-button">Add</button>}
          {(friend.hasOwnProperty('has_sent_request') && friend.has_sent_request === '2') && <button className="request-button">Requested</button>}
          {/* <button className="add-button">Add</button> */}
        </div>
        {/* Other user information can be added here */}
      </div>
    </div>
  );
};

export default FriendCard;