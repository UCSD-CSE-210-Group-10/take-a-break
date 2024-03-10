import React, { useState, useEffect } from "react";
import './Friends.css';
import NavigationBar from './NavigationBar';
import FriendCard from './FriendCard';
import RequestModal from './RequestModal'; 

const Friends = () => {

  const [friends, setFriends] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [foundFriends, setFoundFriends] = useState([]);

  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  const updateSentRequest = (friend) => {

    const updatedFriend = { ...friend, has_sent_request: '2' };
    const friendIndex = foundFriends.findIndex((f) => f.email === friend.email);

    const updatedFoundFriends = [...foundFriends];
    updatedFoundFriends[friendIndex] = updatedFriend;

    setFoundFriends(updatedFoundFriends);
  };


  useEffect(() => {
    const jwtToken = localStorage.getItem('token');
    // Function to fetch friends from the API
    const fetchFriends = async () => {
      try {
        const response = await fetch(`http://localhost:8080/friends/${jwtToken}`);
        const data = await response.json();
        setFriends(data); // Assuming the API response contains an array of friends
      } catch (error) {
        console.error('Error fetching friends:', error);
      }
    };

    // Call the fetchEvents function
    fetchFriends();
  }, []); // Empty dependency array ensures the effect runs once when the component mounts


  

  const handleSearch = async (event) => {
    const jwtToken = localStorage.getItem('token');
    const term = event.target.value;
    setSearchTerm(term);
    if(term.length > 0){
      try {
        const response = await fetch(`http://localhost:8080/friends/search/${jwtToken}?searchTerm=${term}`);
        if (!response.ok) {
          throw new Error('Failed to fetch data');
        }
        const data = await response.json();
        setFoundFriends(data);
      } catch (error) {
        console.error('Error searching friends:', error);
      }
    }
    else {
        setFoundFriends([])
    }
  };

  return (
    <div>
      <NavigationBar />
      <div className="friends-container" data-testid="friends-container">
      <h2 className="content">Discover</h2>
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar-friends">
            <input type="text" placeholder="Search Friend" value={searchTerm} onChange={handleSearch} className="search-input" />
            <div><button className="RequestsButton" onClick={openModal}>Requests</button></div>
          </div>
          <div>
          </div>
        </div>
        <div className="friend-card-container">
          {foundFriends && foundFriends.map((friend) => (
            <FriendCard key={friend.id} friend={friend} updateSentRequest={updateSentRequest}/>
          ))}
        </div>
        <hr/>
        <h2 className="content">My Friends</h2>
        <div className="friend-card-container">
          {friends && friends.map((friend) => (
            <FriendCard key={friend.id} friend={friend} />
          ))}
        </div>
        
        <RequestModal
        isOpen={isModalOpen}
        onRequestClose={closeModal}
        jwtToken = {localStorage.getItem('token')}
      />

      </div>
    </div>
  );
}

export default Friends;
