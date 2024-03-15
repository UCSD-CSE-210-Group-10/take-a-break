import { render, screen } from '@testing-library/react';
import EventsPage from '../events/EventsPage';
import { MemoryRouter } from "react-router-dom";

test("Search Bar Renders Successfully", () => {
  render(<MemoryRouter><EventsPage/></MemoryRouter>);
  
  const searchBar = screen.getByPlaceholderText('Search Event');
  expect(searchBar).toBeInTheDocument();

  const tagsDropdown = screen.getByText('Filter');
  expect(tagsDropdown).toBeInTheDocument();
});

test('Event cards with associated details', async () => {

    const mockEvents = [
      { id: 1, title: 'Event 1', date: "2024-03-22T00:00:00Z", time: "2024-03-22T00:00:00Z", host: 'Host 1' },
      { id: 2, title: 'Event 2', date: "2024-03-22T00:00:00Z", time: "2024-03-22T00:00:00Z", host: 'Host 2' }
    ];


    jest.spyOn(global, 'fetch').mockResolvedValue({
      json: jest.fn().mockResolvedValue(mockEvents),
    });

    render(<MemoryRouter><EventsPage/></MemoryRouter>);


    const eventCards = await screen.findAllByRole('link', { name: /Event \d/ });
    expect(eventCards).toHaveLength(2);

    eventCards.forEach((card, index) => {
        const event = mockEvents[index];
        expect(card).toHaveTextContent(event.title);
        expect(card).toHaveTextContent(event.host);
        expect(card).toHaveTextContent(new Date(event.date.substring(0, event.date.length-1).toLocaleString('en-US')).toDateString());
        expect(card).toHaveTextContent(new Date(event.time.substring(0, event.time.length-1)).toLocaleTimeString("en-US"));
      });
});
