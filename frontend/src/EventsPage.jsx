import React, { useState, useEffect } from "react";
import './EventsPage.css';
import { Link } from "react-router-dom";
import NavigationBar from './NavigationBar';
import EventCard from './EventCard'; // Import the EventCard component

const EventsPage = () => {
  const [events, setEvents] = useState([]);
  const [selectedTags, setSelectedTags] = useState([]);   // State to store selected tags

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

  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [noResultsMessage, setNoResultsMessage] = useState(false);


  const handleSearch = async (event) => {
    const term = event.target.value;
    setSearchTerm(term);
  
    try {
      const response = await fetch(`http://localhost:8080/events/search?searchTerm=${term}`);
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

  const handleTag = (event) => {
    
    // Implementation for searching events
    console.log(event.target.value);

    const selectedOptions = Array.from(event.target.selectedOptions).map(option => option.value);
    setSelectedTags(selectedOptions); // Update selected tags state
  };

  // Function to filter events based on selected tags
  const filteredEvents = selectedTags.length
    ? events.filter(event => selectedTags.every(tag => event.tags.includes(tag)))
    : events;

  return (
    <div>
      <NavigationBar />
      <div className="events-container">
        <div className="content" style={{ backgroundColor: '#FCE7A2' }}>
          <div className="search-bar">
            <input type="text" placeholder="Search Event" value={searchTerm} onChange={handleSearch} className="search-input" />
            <select className="tags-dropdown" multiple onChange={handleTag}>
              <option value=""> Tags </option>
              <option value="Tag1">Tag 1</option>
              <option value="Tag2">Tag 2</option>
              <option value="Tag3">Tag 3</option>
              <option value="Tag4">Tag 4</option>
              <option value="Tag5">Tag 5</option>
              {/* <option value="Physical Wellness">Physical Wellness</option>
              <option value="Cultural Exchange">Cultural Exchange</option>
              <option value="LGBTQ+">LGBTQ+</option>
              <option value="Arts Entertainment">Arts/Entertainment</option>
              <option value="Graduate">Graduate</option>
              <option value="Undergraduate">Undergraduate</option>
              <option value="Virtual">Virtual</option>
              <option value="In Person">In Person</option>
              <option value="Free Food">Free Food</option> */}
            </select>
          </div>
        </div>
        {/* <div className="event-cards">
          {filteredEvents.map(event => (
            <Link key={event.id} to={`/events/${event.id}`} className="event-card-link">
              <div className="event-card">
                <img src={logo} alt="Event" className="event-image" />
                <h3>{event.title}</h3>
                <p>
                  <span>{new Date(event.date).toDateString()}</span> | <span>{new Date(event.time).toLocaleTimeString("en-US")}</span> 
                </p>
                <p>{event.host}</p>
              </div>
            </Link>

          ))}
        </div> */}
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