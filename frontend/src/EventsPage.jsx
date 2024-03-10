import React, { useState, useEffect } from "react";
import './EventsPage.css';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard'; // Import the EventCard component

const EventsPage = () => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await fetch(`${process.env.REACT_APP_SERVER_URL}/events`);
        const data = await response.json();
        setEvents(data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    fetchEvents();
  }, []);

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
      const response = await fetch(`${process.env.REACT_APP_SERVER_URL}/events/search?searchTerm=${term}`);
      if (!response.ok) {
        throw new Error('Failed to fetch search results');
      }
      const data = await response.json();
      setSearchResults(data);
      // Set the message state based on search results
      setNoResultsMessage(term !== '' && data.length === 0);
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
      <NavigationBar />
      <div className="events-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar">
            <input type="text" placeholder="Search Event" value={searchTerm} onChange={handleSearch} className="search-input" />
            <div className="tags-dropdown-container" onClick={toggleDropdown}>
              Tags
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