# Database Documentation

We use PostgreSQL relational database for our project. Since our project involves associating users with events, a relational database made sense, since it will make joins faster and easier. 

Now, we will describe the schema of our database in detail. We have the following tables:

## Events 
The `events` table is designed to store information about various event. It provides a structured representation of essential details related to an event. The table schema is as follows:

- **id (Serial):** Unique identifier for each event. (Primary key, auto-incremented)
- **title (VARCHAR(255)):** Event title. (Not null)
- **venue (VARCHAR(255)):** Location of the event. (Not null)
- **date (DATE):** Date of the event. (Not null)
- **time (TIME):** Start time of the event. (Not null)
- **description (TEXT):** Event details.
- **tags (VARCHAR(255)):** Keywords for categorization.
- **imagepath (VARCHAR(255)):** Path or URL to event image.
- **host (VARCHAR(255)):** Organizer or hosting entity.
- **contact (VARCHAR(255)):** Contact information.


## Users
The `users` table stores the information related to the users. The schema is as follows:
- **email_id (VARCHAR(255)):** Unique email identifier for each user. (Primary key)
- **name (VARCHAR(255)):** User's name. (Not null)
- **role (user_role):** User's role, defined by the `user_role` enumeration. (Not null). The `user_role` enumeration contains two values: *admin* and *user*.

## User Event
The `user_event` table stores the association between users and events. For example if a user A is attending event B, then this table will contain an entry representing that. 

The schema is as follows:

- **id (Serial):** Unique identifier for each user-event association. (Primary key, auto-incremented)
- **email_id (VARCHAR(255)):** Email identifier of the user associated with the event. (Not null). It is a foreign key to the `users` table
- **event_id (INT):** Identifier of the event associated with the user. (Not null). It is a foreign key to the `events` table.


## Friends
The `friends` table stores the association between two users. For example, if a user A is friends with user B, then this table will contain a corresponding entry. Note that the friend association is bidirectional. If (A, B) is present in the table, then (B, A) will also be present in the table. The schema is as follows:

- **id (Serial):** Unique identifier for each friendship. (Primary key, auto-incremented)
- **email_id1 (VARCHAR(255)):** Email identifier of the first user in the friendship. (Not null). It is a foreign key to the `users` table.
- **email_id2 (VARCHAR(255)):** Email identifier of the second user in the friendship. (Not null). It is a foreign key to the `users` table.

## Friend requests
The `friend_requests` table stores the friend requests made. We have the following fields in this table:

- **id (Serial):** Unique identifier for each friend request. (Primary key, auto-incremented)
- **sender (VARCHAR(255)):** Email identifier of the user sending the friend request. (Not null). It is a foreign key to the `users` table.
- **receiver (VARCHAR(255)):** Email identifier of the user receiving the friend request. (Not null). It is a foreign key to the `users` table. 
- **ignored (BOOLEAN):** Indicates whether the friend request has been ignored (default is false). We keep this field since it allows us to prevent an adversary from sending repeated requests to someone who is ignoring them. 
