import { render, screen } from '@testing-library/react';
import EventsPage from '../EventsPage';
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
      { id: 1, title: 'Event 1', date: new Date(), time: new Date(), host: 'Host 1' },
      { id: 2, title: 'Event 2', date: new Date(), time: new Date(), host: 'Host 2' }
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
        expect(card).toHaveTextContent(event.date.toDateString());
        expect(card).toHaveTextContent(event.time.toLocaleTimeString("en-US"));
      });
});
