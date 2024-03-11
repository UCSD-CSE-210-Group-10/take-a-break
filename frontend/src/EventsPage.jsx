import React, { useState, useEffect } from "react";
import './EventsPage.css';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard'; // Import the EventCard component

const EventsPage = ({ handleLogout }) => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const { hostname, protocol } = window.location;
        const response = await fetch(`${protocol}//${hostname}:8080/events`);
        const data = await response.json();
        if(data.error && data.error === "Auth Error") {
          handleLogout()
        }
        setEvents(data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    fetchEvents();
  }, [handleLogout]);

  const toggleDropdown = (event) => {
    setIsDropdownOpen(!isDropdownOpen);
    event.stopPropagation(); 
  };

  const handleTagSelect = (tag, event) => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter(t => t !== tag));
    } else {
      setSelectedTags([...selectedTags, tag]);
    }
    event.stopPropagation(); 
  };

  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [noResultsMessage, setNoResultsMessage] = useState(false);


  const handleSearch = async (event) => {
    const term = event.target.value;
    setSearchTerm(term);
  
    try {
      const { hostname, protocol } = window.location;
      const response = await fetch(`${protocol}//${hostname}:8080/events/search?searchTerm=${term}`);
      if (!response.ok) {
        throw new Error('Failed to fetch search results');
      }
      const data = await response.json();
      // setSearchResults(data);
      // console.log(data)
      // // Set the message state based on search results
      // setNoResultsMessage(term !== '' && data.length === 0);
          // Check if data is not null before accessing its length property
    // Check if data is null or has a length of 0
    if (data === null || data.length === 0) {
      console.log('No search results found.');
      setNoResultsMessage(true);
    } else {
      console.log('Search results:', data);
      setSearchResults(data);
      setNoResultsMessage(false);
    }
    } catch (error) {
      console.error('Error searching events:', error);
    }
  };

  const filteredEvents = selectedTags.length
  ? events.filter(event => selectedTags.every(tag => event.tags.includes(tag)))
  : events;

  useEffect(() => {
    const closeDropdown = () => setIsDropdownOpen(false);
    document.addEventListener('click', closeDropdown);
    return () => document.removeEventListener('click', closeDropdown);
  }, []);


  return (
    <div>
      <NavigationBar handleLogout={handleLogout}/>
      <div className="events-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar">
            <input type="text" placeholder="Search Event" value={searchTerm} onChange={handleSearch} className="search-input" />
            <div className="tags-dropdown-container" onClick={toggleDropdown}>
              Filter
              <div className={`tags-dropdown ${isDropdownOpen ? 'open' : ''}`}>
                {['Tag1', 'Tag2', 'Tag3', 'Tag4', 'Tag5'].map(tag => (
                  <div key={tag} onClick={(e) => handleTagSelect(tag, e)} className={`dropdown-option ${selectedTags.includes(tag) ? 'selected' : ''}`}>
                    {tag}
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
        <div className="event-cards">
          {noResultsMessage ? (
            <p>No events found for "{searchTerm}"</p>
          ) : (
            (searchTerm !== '' && searchResults && searchResults.length > 0 ? (
              searchResults.map(event => (
                <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
                  <EventCard event={event} />
                </Link>
              ))
            ) : (
              filteredEvents.map(event => (
                <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
                  <EventCard event={event} />
                </Link>
              ))
            )))
          }
        </div>
      </div>
    </div>
  );
}

export default EventsPage;  