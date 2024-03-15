import { render, screen } from '@testing-library/react';
import WellfarePage from '../events/wellfare_events/WellfarePage';
import { MemoryRouter } from "react-router-dom";

test("Wellfare Page Renders Successfully", () => {
  render(<MemoryRouter><WellfarePage/></MemoryRouter>);
  
  const title = screen.getByText('Health and Wellfare @ UC San Diego');
  expect(title).toBeInTheDocument();

  const subTitle = screen.getByText('UC San Diego is dedicated to supporting the well-being and academic achievements of every student.');
  expect(subTitle).toBeInTheDocument();

  const upcomingEventsTitle = screen.getByText('Upcoming Events');
  expect(upcomingEventsTitle).toBeInTheDocument();
});

test('Event cards with associated details', async () => {
    const mockEvents = [
      { id: 1, title: 'Event 1', date: new Date(), time: new Date(), host: 'Host 1', tags: ['Tag1', 'Tag2']},
      { id: 2, title: 'Event 2', date: new Date(), time: new Date(), host: 'Host 2', tags: ['Physical Wellness', 'Tag4']}
    ];

    jest.spyOn(global, 'fetch').mockResolvedValue({
      json: jest.fn().mockResolvedValue(mockEvents),
    });

    render(<MemoryRouter><WellfarePage/></MemoryRouter>);

    const eventCards = await screen.findAllByRole('link', { name: /Event \d/ });
    expect(eventCards).toHaveLength(1);

    const card = eventCards[0];
    const event = mockEvents[1];
    expect(card).toHaveTextContent(event.title);
    expect(card).toHaveTextContent(event.host);
    expect(card).toHaveTextContent(event.date.toDateString());
    expect(card).toHaveTextContent(event.time.toLocaleTimeString("en-US"));
});
