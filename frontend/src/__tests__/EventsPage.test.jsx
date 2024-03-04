import { render, screen } from '@testing-library/react';
import EventsPage from '../EventsPage';
import { MemoryRouter } from "react-router-dom";

test("Search Bar Renders Successfully", () => {
  render(<MemoryRouter><EventsPage/></MemoryRouter>);
  
  const searchBar = screen.getByPlaceholderText('Search Event');
  expect(searchBar).toBeInTheDocument();

  const tagsDropdown = screen.getByText('Tags');
  expect(tagsDropdown).toBeInTheDocument();
});

test("Event Cards Render Successfully", () => {
    render(<MemoryRouter><EventsPage/></MemoryRouter>);

    const eventCardLinks = screen.getAllByRole('link', { name: /Event \d/ });
    eventCardLinks.forEach(link => {
        expect(link).toBeInTheDocument();
    });
});

test("Event Details Render Correctly",  () => {
    render(<MemoryRouter><EventsPage/></MemoryRouter>);

  //For a single organization
//   const eventName = screen.getByText('Event 1');
//   expect(eventName).toBeInTheDocument();

//   const eventDate = screen.getByText('February 20, 2024');
//   expect(eventDate).toBeInTheDocument();

//   const eventTime = screen.getByText('10:00 AM');
//   expect(eventTime).toBeInTheDocument();

//   const eventOrganization = screen.getByText('Organization A');
//   expect(eventOrganization).toBeInTheDocument();

  //For multiple organizations
    const eventNames = screen.getAllByRole('heading', { level: 3 });
        eventNames.forEach(name => {
            expect(name).toBeInTheDocument();
        });
        
    const eventDates =  screen.findAllByText(/2024/);
    expect(eventDates).toHaveLength(mockEvents.length);

    const eventTimes =  screen.findAllByText(/AM|PM/);
    expect(eventTimes).toHaveLength(mockEvents.length);

    const eventOrganizations = screen.findAllByText(/Organization/);
    expect(eventOrganizations).toHaveLength(mockEvents.length);
});
