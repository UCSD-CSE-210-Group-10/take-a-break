import { render, screen } from '@testing-library/react'
import EventDetails from '../EventDetails';

test("Sections Render Successfully", () => {
    render(<EventDetails/>);
    
    const leftSec = screen.getByTestId('left-section-events');
    expect(leftSec).toBeInTheDocument();

    const rightSec = screen.getByTestId('right-section-events');
    expect(rightSec).toBeInTheDocument();
})

test("Buttons Render Successfully", () => {
    render(<EventDetails />);

    const backButton = screen.getByAltText('Back');
    expect(backButton).toBeInTheDocument();

    const rsvpButton = screen.getByText('RSVP');
    expect(rsvpButton).toBeInTheDocument();
});

test("Labels Render Successfully", () => {
    render(<EventDetails/>);

    const locationLabel = screen.getByText('Location');
    expect(locationLabel).toBeInTheDocument();

    const dateAndTimeLabel = screen.getByText('Date and Time');
    expect(dateAndTimeLabel).toBeInTheDocument();

    const eventFeeLabel = screen.getByText('Event Fee');
    expect(eventFeeLabel).toBeInTheDocument();

    const contactLabel = screen.getByText('Contact');
    expect(contactLabel).toBeInTheDocument();

    const audienceLabel = screen.getByText('Audience');
    expect(audienceLabel).toBeInTheDocument();

    const eventHostLabel = screen.getByText('Event Host/Organization');
    expect(eventHostLabel).toBeInTheDocument();

    const eventCategoryLabel = screen.getByText('Event Category');
    expect(eventCategoryLabel).toBeInTheDocument();
})
