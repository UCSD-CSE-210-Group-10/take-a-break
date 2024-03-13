# Developer Documentation

## Introduction
The project leverages a modular architecture for efficient organization and maintainability. This section outlines the three core components:

- **Frontend:** This component encompasses the application's user interface, primarily built using React.js. The React frontend code resides within the `frontend` directory.

- **Backend:** This component handles the server-side logic of the application. It is implemented using the Go programming language. The Go backend code is located in the `backend` directory.

- **Database:** This component establishes and manages the application's database. It utilizes PostgreSQL as the relational database management system. The files pertaining to PostgreSQL setup are stored in the `sql` directory.

## Backend
The backend directory contains several packages used to construct our application. We provide a brief overview of each package

#### Database Package
This package encapsulates the database interaction logic. It includes functions for setting up the database connection and performing necessary configurations. Database connection details are retrieved from the `.env` file located within the `backend` directory.


#### Events Package
This package manages event creation and retrieval functionalities. It consists of two primary files:
  - `events_handler.go`: This file is responsible for processing network requests directed towards the backend related to events.
  - `events_operations.go`: This file houses the core logic for event management. It contains functions utilized by the handlers defined in events_handler.go.
  - Unit tests: Unit tests for the events package are located within the dedicated `backend/events/tests` directory.

#### Users Package
This package is similar to the `events` package. It handles user creation and related functionalities such as making friends among users. It consists of two files:
  - `users_handler.go`: This file processes the network requests directed towards creation or fetching of users.
  - `users_operations.go`: This file contains the core logic for user management. It contains functions for making friends, fetching friends, and creating users.
  - Unit tests: Unit tests for the users package are located within the dedicated `backend/users/tests` directory.

#### User_event Package
This package managers the association between the user and the event. It contains handlers for inserting the user and the corresponding event into the `user_event` table in the database when the user RSVPs to an event, and handlers for fetching friends that are attending specific event that the user is looking at.
  - `user_event_handler.go`: This file has handlers for posting UserEvent to the database and also getting a UserEvent from the database. `PostUserEvent` function handles the POST request to create a new user event (RSVP). `GetUserEvent` function handles the GET request to retrieve a user event by email ID and event ID
  - `attendance_handler.go`: This file contains the friend attendance handler which handles fetching the specific users that has the friends relationship with the current user and are attending the specific event that the user is looking at. 
  - Unit tests: Unit tests for the users package are located within the dedicated `backend/user_event/tests` directory.

#### Handle_friend.go
This package consists of one file:
- `handle_friend.go`:This file functions for searching friends based on username or name, deleting a friendship, and handling corresponding HTTP requests. The code uses a PostgreSQL database, and the package includes handlers for searching friends and deleting friendships in a web service. The functions are designed to work with the "gin-gonic/gin" framework and utilize JWT authentication for user identification.
- Unit tests: Unit tests for the login package are located within the dedicated `backend/handle_friend/tests` directory.

#### Login Package
This package consists of one file:
  - `login.go`: The Go code in the "login" package implements user authentication with Google OAuth, exchanging an authorization code for an access token, verifying the user's identity, and inserting user information into a database. 
  - Unit tests: Unit tests for the login package are located within the dedicated `backend/login/tests` directory.

#### Models Package
This package consists of one file:
  - `models.go`: This file contains structs which define the data models used within the application, allowing for structured representation and manipulation of event, user, user-event, user request, and configuration data.

#### Utils Package
This package contains utility functions commonly used in a web service application.
- `web_utils.go`: This file contains 4 utility functions - `HandleNotFound`, `HandleInternalServerError`, `HandleBadRequest`, and `SaveUploadedFile`. These utility functions help handle common scenarios such as handling errors, saving uploaded files, and providing appropriate responses to client requests in a web service application built with the Gin framework in Go.

## Frontend

#### Components
- `CreateEvent`
- `EventCard`
- `EventDetails`
- `EventsPage`
- `FriendCard`
- `Friends`
- `Login`
- `NavigationBar`
- `RequestModal`
- `UserProfile`
- `WelfarePage`
