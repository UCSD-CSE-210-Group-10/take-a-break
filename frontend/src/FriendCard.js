import React from 'react';
import './FriendCard.css'; // Create a CSS file for styling

const FriendCard = ( {friend} ) => {
  return (
    <div className="friend-card">
      <div className="friend-image">
        <img src={friend.image} alt={friend.name} />
      </div>
      <div className="friend-info">
        <div className="friend-header">
          <h3>{friend.name}</h3>
          <button className="add-button">Add</button>
        </div>
        {/* Other user information can be added here */}
      </div>
    </div>
  );
};

export default FriendCard;