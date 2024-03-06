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
  - `users_operations.go`: THis file contains the core logic for user management. It contains functions for making friends, fetching friends, and creating users.
  - Unit tests: Unit tests for the users package are located within the dedicated `backend/users/tests` directory.


 
  
