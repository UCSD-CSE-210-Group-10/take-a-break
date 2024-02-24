CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    venue VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    time TIME NOT NULL,
    description TEXT,
    tags VARCHAR(255),
    imagepath VARCHAR(255),
    host VARCHAR(255),
    contact VARCHAR(255)
);


INSERT INTO events (title, venue, date, time, description, tags, imagepath, host, contact)
VALUES
    ('Event 1', 'Venue 1', '2024-02-17', '18:00', 'Description for Event 1', 'Tag1, Tag2', './images/event1.jpg', 'Host 1', 'Contact 1'),
    ('Event 2', 'Venue 2', '2024-02-18', '19:30', 'Description for Event 2', 'Tag2, Tag3', './images/event2.jpg', 'Host 2', 'Contact 2'),
    ('Event 3', 'Venue 3', '2024-02-19', '20:15', 'Description for Event 3', 'Tag3, Tag4', './images/event3.jpg', 'Host 3', 'Contact 3');
