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

CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS users (
    email_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    avatar TEXT NOT NULL
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
    ignored BOOLEAN DEFAULT false,
    FOREIGN KEY (sender) REFERENCES users(email_id) ON DELETE CASCADE,
    FOREIGN KEY (reciever) REFERENCES users(email_id) ON DELETE CASCADE
);

INSERT INTO events (title, venue, date, time, description, tags, imagepath, host, contact)
VALUES
    (
        'Stress Better Workshop',
        'The Zone',
        '2024-03-22',
        '13:00:00',
        'Every Friday from 1-2 pm, Melissa Hawthorne-Campos, LCSW, from CAPS will lead a Stress Better Workshop in The Zone. Check out the dates below for which topics will be covered from Weeks 7-10:
    Week 7, February 23rd: Don’t Let Your Mind Boss You Around/ Healthy Distractions/ Scheduled Worry Time
    Week 8, March 1st: The Power of Imagination
    Week 9, March 8th: Finding Your Wise Mind/ The Stop Skill
    Week 10, March 15th: Finessing Finals',
        'Graduate, Undergraduate, Physical Wellness, Mental Wellness, In-person',
        'https://d3flpus5evl89n.cloudfront.net/60354997e9d133227dc9f55c/65cfd2451087122ee7d21b7a/scaled_768.jpg',
        'CAPS',
        '858-534-3755'
    ),
    (
        'LGBTQ Community Forum for Graduate and Professional Students',
        'UCSD Cross-Cultural Center Conference Room Price Center East 2nd Floor',
        '2024-03-20',
        '16:00:00',
        'This in-person forum for persons who identify as LGBTQ+ exists to create space for community development and to foster a sense of belonging. The content is guided by community members and often focuses on LGBTQ and/or graduate and professional student issues. Examples include personal and professional relationships, queer topics in the academia and media, intersectionality, minority stress and coping with imposter syndrome. Conversations will be facilitated by Dr David Kersey.
    To learn more, please contact David Kersey MD at dkersey@health.ucsd.edu or 4-3050. Students are welcome to walk-in to a meeting as well.
    Location: UCSD Cross-Cultural Center Conference Room Price Center East 2nd',
        'Graduate, Undergraduate, Physical Wellness, Mental Wellness, LGBTQ+, In-person',
        'https://d3flpus5evl89n.cloudfront.net/60354997e9d133227dc9f55c/658f2716c4854339cc339335/scaled_768.jpg',
        'CAPS, Tritons Flourish (TF)',
        'David Kersey MD at dkersey@health.ucsd.edu or 4-3050'
    ),
    (
        'Cultivating Community',
        'Cross-Cultural Center, 9500 Gilman Dr',
        '2024-02-26',
        '15:00:00',
        'Students, staff, and faculty are welcome to join us for this social justice oriented community building event where we will seek to cultivate community across campus!',
        'Graduate, Undergraduate, Cultural Exchange, In-person',
        'https://d3flpus5evl89n.cloudfront.net/5e8cb604df82fe3367f557ef/65dce11d099b951edece612d/scaled_512.jpg',
        'Cross-Cultural Center at UCSD',
        'cccenter@ucsd.edu'
    ), 
    (
        'Design Co: Escape Room Social',
        'Design & Innovation Building, DIB 208',
        '2024-03-06',
        '9:00:00',
        'Crack the code to unlock the fun!
Take a break from solving the mysteries of your coursework and join us for Design Co’s Escape Room social! Grab your friends and unleash your inner detective to conquer puzzles and mini games to find the secret hex-code. Oh, and did we mention there’s pizza waiting for you at the finish line?
Slice through the end of the quarter with us on Wednesday, 3/13 at 6:30 PM in DIB 208!',
        'Free Food, Graduate, Undergraduate, In-person, Arts / Entertainment',
        'https://d3flpus5evl89n.cloudfront.net/5e8cb604df82fe3367f557ef/65eba035e318f152be017f9d/scaled_512.jpg',
        'Design Co',
        'designatucsd@gmail.com'
    ), 
    (
        'Cafecito Hour',
        'Pepper Canyon Hall 264, conference room',
        '2024-04-10',
        '12:00:00',
        'Join Kimberly Knight-Ortiz, LCSW and your Latinx/Chicanx community Wednesdays from 12-1 weeks 2-10. We will discuss topics impacting Latinx/ Chicanx health, wellbeing and academic success on campus and in the world. This is a space to be in community and uplift one another with collective problem solving, discussion and support. Coffee and light snacks will be provided. Arrive knowing you are welcome exactly as you are. Bilingual dialogue (or even a few words en Español here and there) is welcome if it supports your wellbeing and empowerment. The forum will be in person at Pepper Canyon hall conference room 264. For any questions feel free to contact Kimberly Knight-Ortiz at kknightortiz@health.ucsd.edu',
        'Free Food, Graduate, Undergraduate, Physical Wellness, Mental Wellness, Cultural Exchange, In-person',
        'https://d3flpus5evl89n.cloudfront.net/60354997e9d133227dc9f55c/65f33d8ccb58760cffcc855c/scaled_768.jpg',
        'CAPS, Tritons Flourish (TF)',
        'kknightortiz@health.ucsd.edu'
    ),
    (
        'Gender Buffet and Chill',
        'Women’s Center',
        '2024-03-15',
        '12:00:00',
        'Need to destress before finals? Join us for the last Gender Buffet of the quarter on Friday, March 15 from 12-1pm at the Women’s Center. We will make origami star jars and chill! All genders welcome! Light refreshments will be provided.',
        'Free Food, Graduate, Undergraduate, In-person, Arts / Entertainment',
        'https://d3flpus5evl89n.cloudfront.net/5e8cb604df82fe3367f557ef/65eba0b5e318f152be017fe7/scaled_512.jpg',
        'Women’s Center',
        'women@ucsd.edu'
    )
;


