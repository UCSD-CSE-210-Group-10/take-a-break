import React, { useState, useEffect } from "react";
import './EventsPage.css';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard';

const EventsPage = () => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await fetch('http://localhost:8080/events');
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

  const handleSearch = (event) => {
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
          <div className="search-and-filter">
            <input type="text" placeholder="Search Event" onChange={handleSearch} className="search-input" />
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
          {filteredEvents.map(event => (
            <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
              <EventCard event={event} />
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
}

export default EventsPage;