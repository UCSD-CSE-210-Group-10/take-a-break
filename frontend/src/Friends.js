import React, { useState } from 'react';
import './Friends.css';
import NavigationBar from './NavigationBar';
import FriendCard from './FriendCard';

const Friends = () => {
  // Dummy data for event cards
  // const friends = [
  //   {
  //     id: 1,
  //     name: 'John Doe',
  //     image: './UCSD-logo.png',
  //   },
  //   {
  //       id: 2,
  //       name: 'John Doe',
  //       image: './UCSD-logo.png',
  //   },
  //   {
  //       id: 3,
  //       name: 'John Doe',
  //       image: './UCSD-logo.png',
  //     },

  //   // Add more users as needed
  // ];

  const [searchTerm, setSearchTerm] = useState('');
  const [foundFriends, setFoundFriends] = useState([]);

  const handleSearch = async (event) => {
    const term = event.target.value;
    setSearchTerm(term);
    try {
      const response = await fetch(`http://localhost:8080/search-friends?searchTerm=${term}`);
      if (!response.ok) {
        throw new Error('Failed to fetch data');
      }
      const data = await response.json();
      setFoundFriends(data);
    } catch (error) {
      console.error('Error searching friends:', error);
    }
  };

  return (
    <div>
      <NavigationBar />
      <div className="friends-container" data-testid="friends-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar-friends">
            <input type="text" placeholder="Search Friend" value={searchTerm} onChange={handleSearch} className="search-input" />
          </div>
        </div>
        <div className="friend-card-container">
          {foundFriends && foundFriends.map((friend) => (
            <FriendCard key={friend.id} friend={friend} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default Friends;
