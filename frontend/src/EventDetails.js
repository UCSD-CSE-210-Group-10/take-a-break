function EventDetails({
  eventName,
  location,
  dateTime,
  eventFee,
  contact,
  audience,
  host,
  category,
  posterUrl,
}) {
  return (
    <div className="event-details">
      <div className="poster">
        <img src={posterUrl} alt="Event Poster" />
      </div>
      <div className="details">
        <h2>{eventName}</h2>
        <p>
          <strong>Location:</strong> {location}
        </p>
        <p>
          <strong>Date & Time:</strong> {dateTime}
        </p>
        <p>
          <strong>Fee:</strong> {eventFee}
        </p>
        <p>
          <strong>Contact:</strong> {contact}
        </p>
        <p>
          <strong>Audience:</strong> {audience}
        </p>
        <p>
          <strong>Host/Organization:</strong> {host}
        </p>
        <p>
          <strong>Category:</strong> {category}
        </p>
      </div>
    </div>
  );
}

export default EventDetails;
