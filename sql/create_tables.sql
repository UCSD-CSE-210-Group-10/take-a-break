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

CREATE TABLE IF NOT EXISTS wellfare_events (
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

CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS users (
    email_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    role user_role NOT NULL
);

CREATE TABLE IF NOT EXISTS user_event (
    id SERIAL PRIMARY KEY,
    email_id VARCHAR(255) NOT NULL,
    event_id INT NOT NULL,
    FOREIGN KEY (email_id) REFERENCES users(email_id) ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friends (
    id SERIAL PRIMARY KEY,
    email_id1 VARCHAR(255) NOT NULL,
    email_id2 VARCHAR(255) NOT NULL,
    FOREIGN KEY (email_id1) REFERENCES users(email_id) ON DELETE CASCADE,
    FOREIGN KEY (email_id2) REFERENCES users(email_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friend_requests (
    id SERIAL PRIMARY KEY,
    sender VARCHAR(255) NOT NULL,
    reciever VARCHAR(255) NOT NULL,
    ignored BOOLEAN DEFAULT false
    FOREIGN KEY (sender) REFERENCES users(email_id) ON DELETE CASCADE,
    FOREIGN KEY (reciever) REFERENCES users(email_id) ON DELETE CASCADE
);

INSERT INTO events (title, venue, date, time, description, tags, imagepath, host, contact)
VALUES
    ('Event 1', 'Venue 1', '2024-02-17', '18:00', 'Description for Event 1', 'Tag1, Tag2', './images/event1.jpg', 'Host 1', 'Contact 1'),
    ('Event 2', 'Venue 2', '2024-02-18', '19:30', 'Description for Event 2', 'Tag2, Tag3', './images/event2.jpg', 'Host 2', 'Contact 2'),
    ('Event 3', 'Venue 3', '2024-02-19', '20:15', 'Description for Event 3', 'Tag3, Tag4', './images/event3.jpg', 'Host 3', 'Contact 3');

INSERT INTO wellfare_events (title, venue, date, time, description, tags, imagepath, host, contact)
VALUES
    ('Wellfare Event 1', 'Student Center', '2024-03-10', '15:00', 'A fun gathering for college students to relax and socialize.', 'College, Fun', './images/college_event1.jpg', 'Student Association', 'Contact Student Association'),
    ('Wellfare Event 2', 'Outdoor Amphitheater', '2024-03-15', '16:30', 'An outdoor concert featuring local bands, perfect for unwinding after classes.', 'College, Music', './images/college_event2.jpg', 'Music Club', 'Contact Music Club'),
    ('Wellfare Event 3', 'Library Conference Room', '2024-03-20', '14:00', 'A workshop on stress management and study techniques to help students cope with academic pressure.', 'College, Workshop', './images/college_event3.jpg', 'Counseling Center', 'Contact Counseling Center');

INSERT INTO users (email_id, name, role) VALUES
('admin@example.com', 'Admin User', 'admin'),
('user1@example.com', 'Regular User 1', 'user'),
('user2@example.com', 'Regular User 2', 'user');

INSERT INTO user_event (email_id, event_id) VALUES
('user1@example.com', 1),
('user1@example.com', 2),
('user2@example.com', 2),
('user2@example.com', 3);

INSERT INTO friends (email_id1, email_id2) VALUES
('admin@example.com', 'user1@example.com'),
('user1@example.com', 'admin@example.com'),
('user1@example.com', 'user2@example.com'),
('user2@example.com', 'user1@example.com');

