import React from 'react';
import './FriendCard.css'; // Create a CSS file for stylings

const FriendCard = ( {friend, showButtonType} ) => {

  const sendRequest = async (requestId) => {
    const jwtToken = localStorage.getItem('token');
    // Implement logic to accept the friend request
    try {
      const response = await fetch(`http://localhost:8080/friends/request/send/${jwtToken}`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email_id: `${requestId}`  }),
			});
	
			if (!response.ok) {
				throw new Error("Failed to Send Request");
			}
      
    } catch (error) {
      console.error('Error Sendings request:', error);
    }
  };

  return (
    <div className="friend-card" data-testid="friend-card">
      <div className="friend-image">
        <img src="./UCSD-logo.png" alt={friend.name} />
      </div>
      <div className="friend-info">
        <div className="friend-header">
          <h3>{friend.name}</h3>
          {(showButtonType === '0') && <button onClick={() => sendRequest(friend.email)} className="add-button">Add</button>}
          {(showButtonType === '2') && <button onClick={() => sendRequest(friend.email)} className="request-button">Requested</button>}
          {/* <button className="add-button">Add</button> */}
        </div>
        {/* Other user information can be added here */}
      </div>
    </div>
  );
};

export default FriendCard;