import React, { useState, useEffect } from "react";
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

  const [friends, setFriends] = useState([]);

  useEffect(() => {
    // Function to fetch friends from the API
    const fetchFriends = async () => {
      try {
        const response = await fetch('http://localhost:8080/friends');
        const data = await response.json();
        setFriends(data); // Assuming the API response contains an array of events
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    };

    // Call the fetchEvents function
    fetchFriends();
  }, []); // Empty dependency array ensures the effect runs once when the component mounts


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
