import React from 'react';
import './Friends.css';
import logo from './UCSD-logo.png';
// import eventImage from './event-image.jpg';
import NavigationBar from './NavigationBar';
import FriendCard from './FriendCard';

const Friends = () => {
  // Dummy data for event cards
  const friends = [
    {
      id: 1,
      name: 'John Doe',
      image: './UCSD-logo.png',
    },
    {
        id: 2,
        name: 'John Doe',
        image: './UCSD-logo.png',
    },
    {
        id: 3,
        name: 'John Doe',
        image: './UCSD-logo.png',
      },

    // Add more users as needed
  ];

  // Functionality for searching events
  const handleSearch = (event) => {
    // Implementation for searching events
    console.log(event.target.value);
  };

  return (
    <div>
      <NavigationBar />
      <div className="friends-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar-friends">
            <input type="text" placeholder="Search Event" onChange={handleSearch} className="search-input" />
          </div>
        </div>
        <div className="friend-card-container">
            {friends.map(friend => (
                <FriendCard key={friend.id} friend={friend} />
            ))}
        </div>
      </div>
    </div>
  );
}

export default Friends;
